package cmd

import (
	"fmt"
	"strings"
	"music-management-system/pkg/constants"
	"github.com/AlecAivazis/survey/v2"
	"golang.org/x/exp/slices"
)

type option struct {
	Name  string
	Items []map[int8]string
}

var options = []option{
	{
		Name: "album",
		Items: []map[int8]string{
			{1: "Create new album"},
			{2: "Get all albums"},
			{3: "Get album info by id"},
			{4: "Delete album by id"},
			{5: "Update album by id"},
			{6: "Get all tracks of album"},
		},
	},
	{
		Name: "artist",
		Items: []map[int8]string{
			{7: "Create new artist"},
			{8: "Get all artists"},
			{9: "Get artist by id"},
			{10: "Get all albums of artist"},
			{11: "Get all songs of artist"},
			{12: "Delete artist by id"},
			{13: "Update artist by id"},
		},
	},
	{
		Name: "genre",
		Items: []map[int8]string{
			{14: "Create new genre"},
			{15: "Get all genres"},
			{16: "Get genre by id"},
			{17: "Delete genre by id"},
			{18: "Update genre by id"},
			{19: "Get all tracks of genre"},
		},
	},
	{
		Name: "track",
		Items: []map[int8]string{
			{20: "Create new track"},
			{21: "Get all tracks"},
			{22: "Get track by id"},
			{23: "Delete track by id"},
			{24: "Update track by id"},
		},
	},
	{
		Name: "playlist",
		Items: []map[int8]string{
			{25: "Create new playlist"},
			{26: "Get all playlists"},
			{27: "Get playlist by id"},
			{28: "Delete playlist by id"},
			{29: "Update playlist by id"},
			
			{30: "Add track to playlist"},
			{31: "Get all tracks of playlist"},
			{32: "Remove track from playlist"},
			{33: "Get all playlists have track"},
		},
	},
	{
		Name: "user",
		Items: []map[int8]string{
			{34: "Create new user"},
			{35: "Get all users"},
			{36: "Get user by id"},
			{37: "Delete user by id"},
			{38: "Update user by id"},
		},
	},
}


func (d *Delivery) Run() int8{
	var scopes []string

	for _, option := range options {
		scopes = append(scopes, fmt.Sprintf(" ⦿ %v", option.Name))
	}

	var qs = []*survey.Question{
		{
			Name: "scope",
			Prompt: &survey.Select{
				Message: "Choose a scope that you want:",
				Options: scopes,
			},
		},
	}

	answers := struct {
        Scopes string `survey:"scope"`
    }{}

    err := survey.Ask(qs, &answers)
    if err != nil {
        fmt.Println(err.Error())
		return -1
    }

	type Choice struct {
		Key int8
		Value string
	}
	
	var choices []Choice
	var choiceValue []string

	for _, option := range options {
		if strings.Contains(answers.Scopes, option.Name) {
			for _, item := range option.Items {
				for k, v := range item {
					choiceValue = append(choiceValue, fmt.Sprintf(constants.CHOICE_ICON+v))
					getResult := &Choice{
						Key:   k,
						Value: v,
					}
					choices = append(choices, *getResult)
				}
			}
			break
		}
	}

	qs2 := []*survey.Question{
		{
			Name: "options",
			Prompt: &survey.Select{
				Message: "Choose a options:",
				Options: choiceValue,
			},
		},
	}

	answers2 := struct {
        Options string `survey:"options"`
    }{}

    err2 := survey.Ask(qs2, &answers2)
    if err2 != nil {
        fmt.Println(err2.Error())
		return -1
    }

	idx := slices.IndexFunc(choices, func(c Choice) bool {
		return strings.Contains(answers2.Options, c.Value) 
	})

	return int8(choices[idx].Key)
}
