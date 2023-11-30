package verification

import "gorm.io/gorm"

type VerificationCode struct {
	gorm.Model
	Email       *string `json:"email" gorm:"index:verification_email_index"`
	AreaCode    *string `json:"areaCode" gorm:"type:CHAR(8);index:verification_area_code_index"`
	PhoneNumber *string `json:"phonenumber" gorm:"index:verification_phone_number_index"`
	Code        string  `json:"code" gorm:"type:VARCHAR(12)"`
	Purpose     *string `json:"purpost" gorm:"type:VARCHAR(12)"`
}
