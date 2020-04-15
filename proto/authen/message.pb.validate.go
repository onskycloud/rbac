// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: authen/message.proto

package message

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _message_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on EmptyRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *EmptyRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// EmptyRequestValidationError is the validation error returned by
// EmptyRequest.Validate if the designated constraints aren't met.
type EmptyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyRequestValidationError) ErrorName() string { return "EmptyRequestValidationError" }

// Error satisfies the builtin error interface
func (e EmptyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmptyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyRequestValidationError{}

// Validate checks the field values on CustomerRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CustomerRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CustomerNumber

	return nil
}

// CustomerRequestValidationError is the validation error returned by
// CustomerRequest.Validate if the designated constraints aren't met.
type CustomerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CustomerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CustomerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CustomerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CustomerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CustomerRequestValidationError) ErrorName() string { return "CustomerRequestValidationError" }

// Error satisfies the builtin error interface
func (e CustomerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCustomerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CustomerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CustomerRequestValidationError{}

// Validate checks the field values on CustomerByUserNameRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CustomerByUserNameRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UserName

	return nil
}

// CustomerByUserNameRequestValidationError is the validation error returned by
// CustomerByUserNameRequest.Validate if the designated constraints aren't met.
type CustomerByUserNameRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CustomerByUserNameRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CustomerByUserNameRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CustomerByUserNameRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CustomerByUserNameRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CustomerByUserNameRequestValidationError) ErrorName() string {
	return "CustomerByUserNameRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CustomerByUserNameRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCustomerByUserNameRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CustomerByUserNameRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CustomerByUserNameRequestValidationError{}

// Validate checks the field values on CustomerResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CustomerResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CustomerNumber

	// no validation rules for Alias

	// no validation rules for Mobile

	// no validation rules for FirstName

	// no validation rules for LastName

	// no validation rules for Address1

	// no validation rules for Address2

	// no validation rules for Email

	if v, ok := interface{}(m.GetSetting()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CustomerResponseValidationError{
				field:  "Setting",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CustomerResponseValidationError is the validation error returned by
// CustomerResponse.Validate if the designated constraints aren't met.
type CustomerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CustomerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CustomerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CustomerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CustomerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CustomerResponseValidationError) ErrorName() string { return "CustomerResponseValidationError" }

// Error satisfies the builtin error interface
func (e CustomerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCustomerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CustomerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CustomerResponseValidationError{}

// Validate checks the field values on CustomerSetting with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CustomerSetting) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for IsAllowRegister

	if v, ok := interface{}(m.GetSmtp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CustomerSettingValidationError{
				field:  "Smtp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetPayment()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CustomerSettingValidationError{
				field:  "Payment",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Other

	return nil
}

// CustomerSettingValidationError is the validation error returned by
// CustomerSetting.Validate if the designated constraints aren't met.
type CustomerSettingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CustomerSettingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CustomerSettingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CustomerSettingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CustomerSettingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CustomerSettingValidationError) ErrorName() string { return "CustomerSettingValidationError" }

// Error satisfies the builtin error interface
func (e CustomerSettingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCustomerSetting.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CustomerSettingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CustomerSettingValidationError{}

// Validate checks the field values on SMTPConfig with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SMTPConfig) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Host

	// no validation rules for Port

	// no validation rules for Password

	// no validation rules for Account

	return nil
}

// SMTPConfigValidationError is the validation error returned by
// SMTPConfig.Validate if the designated constraints aren't met.
type SMTPConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SMTPConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SMTPConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SMTPConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SMTPConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SMTPConfigValidationError) ErrorName() string { return "SMTPConfigValidationError" }

// Error satisfies the builtin error interface
func (e SMTPConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSMTPConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SMTPConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SMTPConfigValidationError{}

// Validate checks the field values on PaymentConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PaymentConfig) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// PaymentConfigValidationError is the validation error returned by
// PaymentConfig.Validate if the designated constraints aren't met.
type PaymentConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PaymentConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PaymentConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PaymentConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PaymentConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PaymentConfigValidationError) ErrorName() string { return "PaymentConfigValidationError" }

// Error satisfies the builtin error interface
func (e PaymentConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPaymentConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PaymentConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PaymentConfigValidationError{}

// Validate checks the field values on VerifyPermissionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *VerifyPermissionRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Action

	return nil
}

// VerifyPermissionRequestValidationError is the validation error returned by
// VerifyPermissionRequest.Validate if the designated constraints aren't met.
type VerifyPermissionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyPermissionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyPermissionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyPermissionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyPermissionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyPermissionRequestValidationError) ErrorName() string {
	return "VerifyPermissionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e VerifyPermissionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyPermissionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyPermissionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyPermissionRequestValidationError{}

// Validate checks the field values on VerifyPermissionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *VerifyPermissionResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Success

	if v, ok := interface{}(m.GetAuthUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return VerifyPermissionResponseValidationError{
				field:  "AuthUser",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// VerifyPermissionResponseValidationError is the validation error returned by
// VerifyPermissionResponse.Validate if the designated constraints aren't met.
type VerifyPermissionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyPermissionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyPermissionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyPermissionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyPermissionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyPermissionResponseValidationError) ErrorName() string {
	return "VerifyPermissionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e VerifyPermissionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyPermissionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyPermissionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyPermissionResponseValidationError{}

// Validate checks the field values on AuthUser with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *AuthUser) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for CustomerID

	// no validation rules for CustomerNumber

	// no validation rules for Username

	// no validation rules for UserUuid

	return nil
}

// AuthUserValidationError is the validation error returned by
// AuthUser.Validate if the designated constraints aren't met.
type AuthUserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthUserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthUserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthUserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthUserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthUserValidationError) ErrorName() string { return "AuthUserValidationError" }

// Error satisfies the builtin error interface
func (e AuthUserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthUserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthUserValidationError{}

// Validate checks the field values on DeviceTokenResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeviceTokenResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UserId

	// no validation rules for Ip

	// no validation rules for Location

	// no validation rules for Lat

	// no validation rules for Long

	// no validation rules for Active

	// no validation rules for IsMainDevice

	// no validation rules for DeviceToken

	// no validation rules for DeviceType

	return nil
}

// DeviceTokenResponseValidationError is the validation error returned by
// DeviceTokenResponse.Validate if the designated constraints aren't met.
type DeviceTokenResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeviceTokenResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeviceTokenResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeviceTokenResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeviceTokenResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeviceTokenResponseValidationError) ErrorName() string {
	return "DeviceTokenResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeviceTokenResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeviceTokenResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeviceTokenResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeviceTokenResponseValidationError{}
