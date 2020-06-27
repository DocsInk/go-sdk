/*
 * DocsInk REST API
 *
 * DocsInk REST API
 *
 * API version: 1.0.0
 */

package api

// ClientUser struct for ClientUser
type ClientUser struct {
	Active                      float32 `json:"active,omitempty"`
	ChatStatus                  string  `json:"chat_status,omitempty"`
	ConsoleTourbusVerify        float32 `json:"console_tourbus_verify,omitempty"`
	Email                       string  `json:"email,omitempty"`
	EmailAppointmentSchedule    float32 `json:"email_appointment_schedule,omitempty"`
	FirstName                   string  `json:"first_name,omitempty"`
	Image                       string  `json:"image,omitempty"`
	IsInterfaceUser             float32 `json:"is_interface_user,omitempty"`
	LanguageId                  float32 `json:"language_id,omitempty"`
	LastName                    string  `json:"last_name,omitempty"`
	LastSeenDatetime            string  `json:"last_seen_datetime,omitempty"`
	MessageRefreshInterval      float32 `json:"message_refresh_interval,omitempty"`
	OrgUuid                     float32 `json:"org_uuid,omitempty"`
	OrgVerified                 float32 `json:"org_verified,omitempty"`
	PhoneNumber                 string  `json:"phone_number,omitempty"`
	ReceivePushNotices          float32 `json:"receive_push_notices,omitempty"`
	TelehealthAvailability      float32 `json:"telehealth_availability,omitempty"`
	Teleheath5MinuteCheckIn     bool    `json:"teleheath_5_minute_check_in,omitempty"`
	Teleheath5MinuteCheckInType string  `json:"teleheath_5_minute_check_in_type,omitempty"`
	TermsAccepted               bool    `json:"terms_accepted,omitempty"`
	UserType                    string  `json:"user_type,omitempty"`
	Uuid                        float32 `json:"uuid,omitempty"`
	Verified                    float32 `json:"verified,omitempty"`
}
