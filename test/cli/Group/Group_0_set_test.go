package group_test

import (
	"testing"

	cliTest "github.com/Telmate/proxmox-api-go/test/cli"
	"github.com/Telmate/proxmox-api-go/test/cli/Group/group_sub_tests"
	"github.com/Telmate/proxmox-api-go/test/data/test_data_cli"
)

// Create group with all option populated
// Check if populated
// Update group with --members not defined (should not update memberships)
// Check no changes
// Update group with all options empty
// Check empty
// Delete items
// Check items deleted

func Test_Group_0_Cleanup(t *testing.T) {
	// remove group
	Test := &cliTest.Test{
		ReqErr:      true,
		ErrContains: "group0",
		Args:        []string{"-i", "delete", "group", "group0"},
	}
	Test.StandardTest(t)
	// remove user 00
	Test = &cliTest.Test{
		ReqErr:      true,
		ErrContains: "group0-user00@pve",
		Args:        []string{"-i", "delete", "user", "group0-user00@pve"},
	}
	Test.StandardTest(t)
}

func Test_Group_0_Create_User(t *testing.T) {
	Test := &cliTest.Test{
		Contains: []string{"group0-user00@pve"},
		Args:     []string{"-i", "set", "user", "group0-user00@pve"},
	}
	Test.StandardTest(t)
}

func Test_Group_0_Set_Full_Create(t *testing.T) {
	Test := &cliTest.Test{
		Contains: []string{"(group0)"},
		Args:     []string{"-i", "set", "group", "group0", "comment", "--members=root@pam,group0-user00@pve,group0-user01@pve"},
	}
	Test.StandardTest(t)
}

func Test_Group_0_Get_Full_0(t *testing.T) {
	Test := &cliTest.Test{
		NotExpected: "group0-user01",
		NotContains: true,
		Args:        []string{"-i", "get", "group", "group0"},
	}
	out := Test.StandardTest(t)
	group_sub_tests.Get_Test(t, test_data_cli.Group_Get_Full_testData(0), out)
}

func Test_Group_0_Set_MembersNotDefined_Update(t *testing.T) {
	Test := &cliTest.Test{
		Contains: []string{"(group0)"},
		Args:     []string{"-i", "set", "group", "group0", "comment"},
	}
	Test.StandardTest(t)
}

func Test_Group_0_Get_Full_1(t *testing.T) {
	Test := &cliTest.Test{
		Args: []string{"-i", "get", "group", "group0"},
	}
	out := Test.StandardTest(t)
	group_sub_tests.Get_Test(t, test_data_cli.Group_Get_Full_testData(0), out)
}

func Test_Group_0_Set_Empty_Create(t *testing.T) {
	Test := &cliTest.Test{
		Contains: []string{"(group0)"},
		Args:     []string{"-i", "set", "group", "group0", "--members="},
	}
	Test.StandardTest(t)
}

func Test_Group_0_Get_Empty(t *testing.T) {
	Test := &cliTest.Test{
		Args: []string{"-i", "get", "group", "group0"},
	}
	out := Test.StandardTest(t)
	group_sub_tests.Get_Test(t, test_data_cli.Group_Get_Empty_testData(0), out)
}

func Test_Group_0_Delete_Group(t *testing.T) {
	Test := &cliTest.Test{
		Contains: []string{"group0"},
		Args:     []string{"-i", "delete", "group", "group0"},
	}
	Test.StandardTest(t)
}

func Test_Group_0_Delete_User(t *testing.T) {
	Test := &cliTest.Test{
		Contains: []string{"group0-user00@pve"},
		Args:     []string{"-i", "delete", "user", "group0-user00@pve"},
	}
	Test.StandardTest(t)
}

func Test_Group_0_List_Group_NotExistent(t *testing.T) {
	Test := &cliTest.Test{
		NotExpected: "group0",
		NotContains: true,
		Args:        []string{"-i", "list", "groups"},
	}
	Test.StandardTest(t)
}

func Test_Group_0_List_User_NotExistent(t *testing.T) {
	Test := &cliTest.Test{
		NotExpected: "group0-user00@pve",
		NotContains: true,
		Args:        []string{"-i", "list", "users"},
	}
	Test.StandardTest(t)
}
