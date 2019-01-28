package innovaphone

import (
	"sync"
	"testing"
	"time"

	"github.com/regiohelden/innovazammad/config"
)

type numberGroupFixture struct {
	users          map[string][]string
	flipped        bool
	findUserCalled bool
	sync.Mutex
}

var numberGroups = map[string][]string{
	"1": []string{"someothergroup"},
	"3": []string{"somegroup"},
}

func (gc *numberGroupFixture) FindUser(_, _, _, _, cn, _, e164 string, _, _ int) (*FindUserInfoArray, error) {
	gc.Lock()
	defer gc.Unlock()
	gc.findUserCalled = true
	users := &FindUserInfoArray{}
	if fixtureGroups, ok := gc.users[e164]; ok {
		users.Items = make([]*UserInfo, 1)
		groupArray := &GroupArray{Items: make([]*Group, len(fixtureGroups))}
		users.Items[0] = &UserInfo{Groups: groupArray}
		for i, group := range gc.users[e164] {
			groupArray.Items[i] = &Group{Group: group}
		}

	}
	return users, nil
}

func (gc *numberGroupFixture) IsDirectionFlipped() bool {
	return gc.flipped
}

func TestCallInSession_ShouldHandle(t *testing.T) {
	dur := config.Duration(60 * time.Second) // give it enough time to run tests
	config.Global.PBX.GroupCacheTime = &dur

	type fields struct {
		sessionInterface sessionInterface
		CallInfo         *CallInfo
	}
	tests := []struct {
		name          string
		filterOnGroup string
		fields        fields
		want          bool
		shouldCallAPI bool // since the cache is global, this is sensitive to the order in which the tests are called!
	}{
		{"do not filter",
			"",
			fields{
				sessionInterface: &numberGroupFixture{users: numberGroups},
				CallInfo:         &CallInfo{},
			},
			true, false},
		{"filter and not find",
			"somegroup",
			fields{
				sessionInterface: &numberGroupFixture{users: numberGroups},
				CallInfo: &CallInfo{
					No: &NoArray{
						Items: []*No{
							&No{Cn: "somecaller", E164: "1", Type: "peer"},
							&No{Cn: "somecallee", E164: "2", Type: "this"},
						},
					},
				},
			},
			false, true},
		{"filter and find peer",
			"somegroup",
			fields{
				sessionInterface: &numberGroupFixture{users: numberGroups},
				CallInfo: &CallInfo{
					No: &NoArray{
						Items: []*No{
							&No{Cn: "somecaller", E164: "3", Type: "peer"},
							&No{Cn: "somecallee", E164: "4", Type: "this"},
						},
					},
				},
			},
			true, true},
		{"filter (cached) and not find",
			"somegroup",
			fields{
				sessionInterface: &numberGroupFixture{users: numberGroups},
				CallInfo: &CallInfo{
					No: &NoArray{
						Items: []*No{
							&No{Cn: "somecaller", E164: "1", Type: "peer"},
							&No{Cn: "somecallee", E164: "2", Type: "this"},
						},
					},
				},
			},
			false, false},
		{"filter (cached) and not find",
			"somegroup",
			fields{
				sessionInterface: &numberGroupFixture{users: numberGroups},
				CallInfo: &CallInfo{
					No: &NoArray{
						Items: []*No{
							&No{Cn: "somecaller", E164: "3", Type: "peer"},
							&No{Cn: "somecallee", E164: "4", Type: "this"},
						},
					},
				},
			},
			true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.Global.PBX.FilterOnGroup = tt.filterOnGroup
			call := &CallInSession{
				sessionInterface: tt.fields.sessionInterface,
				CallInfo:         tt.fields.CallInfo,
			}
			if got := call.ShouldHandle(); got != tt.want {
				t.Errorf("CallInSession.ShouldHandle() = %v, want %v", got, tt.want)
			}

			// ugly time-sensitive cache access testing
			if apiCalled := tt.fields.sessionInterface.(*numberGroupFixture).findUserCalled; apiCalled != tt.shouldCallAPI {
				t.Errorf("CallInSession.ShouldHandle() API call = %v, want %v", apiCalled, tt.shouldCallAPI)
			}
		})
	}
}
