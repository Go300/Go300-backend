package model

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateMember(t *testing.T) {
	Convey("Member", t, func() {
		member := Member{
			Username: "MuslimBeibytuly",
		}
		newMember, err := CreateMember(member)
		So(len(newMember.Username), ShouldEqual, 15)
		So(len(newMember.Id), ShouldEqual, 12)
		So(err, ShouldEqual, nil)
	})
}

func TestCreateMember2(t *testing.T) {
	Convey("Empty Member Username", t, func() {
		member := Member{
			Username: "",
		}
		_, err := CreateMember(member)
		So(err, ShouldNotEqual, nil)
	})
}
