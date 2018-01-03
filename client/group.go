package client

import (
	"fmt"
)

type GroupsType struct {
	Groups     []string `json:"groups,omitempty"  structs:"groups,omitempty`
	Message      string `json:"message,omitempty" structs:"message,omitempty"`
	Status       string `json:"status,omitempty" structs:"status,omitempty"`
}

type MembersType struct {
	Users      []map[string]string `json:"users,omitempty"  structs:"users,omitempty`
	Status       string `json:"status,omitempty" structs:"status,omitempty"`
}

type AddGroupsType struct {
	Groups      []string `json:"groups,omitempty"  structs:"groups,omitempty`
}

type AddGroupsResponseType struct {
	GroupsAdded   []string `json:"groupsAdded,omitempty"  structs:"groupsAdded,omitempty`
	GroupsSkipped []string `json:"groupsSkipped,omitempty"  structs:"groupsSkipped,omitempty`
	Message         string `json:"message,omitempty" structs:"message,omitempty"`
	Status          string `json:"status,omitempty" structs:"status,omitempty"`
}

type AddUsersType struct {
	Users      []string `json:"users,omitempty"  structs:"users,omitempty`
}

func  (c *ConfluenceClient) GetGroups() (*GroupsType) {
	var u string
	u = fmt.Sprintf("/rest/extender/1.0/group/getGroups")

	groups := new(GroupsType)
	c.debug=true
	res := c.doRequest("GET", u , nil, &groups)

	fmt.Println("res: " + string(res))

	return groups
}

func  (c *ConfluenceClient) GetGroupMembers(groupname string ) (*MembersType) {
	var u string
	u = fmt.Sprintf("/rest/extender/1.0/group/getUsers/" +groupname)

	members := new(MembersType)
	c.debug=true
	res := c.doRequest("GET", u , nil, &members)

	fmt.Println("res: " + string(res))

	return members
}

func  (c *ConfluenceClient) AddGroup(groupname string ) (*AddGroupsResponseType) {
	var u string
	u = fmt.Sprintf("/rest/extender/1.0/group/addGroups")

	var payload = new (AddGroupsType)
	payload.Groups = append(payload.Groups, groupname)

	groups := new(AddGroupsResponseType)
	c.debug=true
	res := c.doRequest("POST", u , payload, &groups)

	fmt.Println("res: " + string(res))

	return groups
}


func  (c *ConfluenceClient) AddGroupMember(groupname string, member string ) (*AddGroupsResponseType) {
	var u string
	u = fmt.Sprintf("/rest/extender/1.0/group/addUsers/" +groupname)

	var payload = new (AddUsersType)
	payload.Users = append(payload.Users, member)

	members := new(AddGroupsResponseType)
	c.debug=true
	res := c.doRequest("POST", u , payload, &members)

	fmt.Println("res: " + string(res))

	return members
}
