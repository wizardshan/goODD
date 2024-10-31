// Code generated by ent, DO NOT EDIT.

package ent

import (
	"goODD/chapter7/repository/ent/schema"
	"goODD/chapter7/repository/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescHashID is the schema descriptor for hash_id field.
	userDescHashID := userFields[1].Descriptor()
	// user.DefaultHashID holds the default value on creation for the hash_id field.
	user.DefaultHashID = userDescHashID.Default.(string)
	// userDescMobile is the schema descriptor for mobile field.
	userDescMobile := userFields[2].Descriptor()
	// user.DefaultMobile holds the default value on creation for the mobile field.
	user.DefaultMobile = userDescMobile.Default.(string)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[3].Descriptor()
	// user.DefaultPassword holds the default value on creation for the password field.
	user.DefaultPassword = userDescPassword.Default.(string)
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[4].Descriptor()
	// user.DefaultAge holds the default value on creation for the age field.
	user.DefaultAge = userDescAge.Default.(int64)
	// userDescLevel is the schema descriptor for level field.
	userDescLevel := userFields[5].Descriptor()
	// user.DefaultLevel holds the default value on creation for the level field.
	user.DefaultLevel = userDescLevel.Default.(int64)
	// userDescNickname is the schema descriptor for nickname field.
	userDescNickname := userFields[6].Descriptor()
	// user.DefaultNickname holds the default value on creation for the nickname field.
	user.DefaultNickname = userDescNickname.Default.(string)
	// userDescAvatar is the schema descriptor for avatar field.
	userDescAvatar := userFields[7].Descriptor()
	// user.DefaultAvatar holds the default value on creation for the avatar field.
	user.DefaultAvatar = userDescAvatar.Default.(string)
	// userDescBio is the schema descriptor for bio field.
	userDescBio := userFields[8].Descriptor()
	// user.DefaultBio holds the default value on creation for the bio field.
	user.DefaultBio = userDescBio.Default.(string)
	// userDescAmount is the schema descriptor for amount field.
	userDescAmount := userFields[9].Descriptor()
	// user.DefaultAmount holds the default value on creation for the amount field.
	user.DefaultAmount = userDescAmount.Default.(int64)
	// userDescStatus is the schema descriptor for status field.
	userDescStatus := userFields[10].Descriptor()
	// user.DefaultStatus holds the default value on creation for the status field.
	user.DefaultStatus = userDescStatus.Default.(int64)
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userFields[11].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userFields[12].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
}
