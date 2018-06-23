package model

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTrivial(t *testing.T) {
	Convey("Member", t, func() {
		member := Member{
			Name:  "test name 1",
			Token: "",
		}
		newMember, err := CreateMember(member)
		So(len(newMember.Token), ShouldEqual, 24)
		So(err, ShouldEqual, nil)
	})
}

func TestEmptyUsername(t *testing.T) {
	Convey("Username", t, func() {
		member := Member{
			Name:  "",
			Token: "",
		}
		_, err := CreateMember(member)
		So(err, ShouldNotEqual, nil)
	})
}
