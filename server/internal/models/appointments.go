package models

import "time"

type Appointment struct {
	Id            string            `json:"id"`
	BusinessId    string            `json:"business_id"`
	TimeSlotId    string            `json:"time_slot_id"`
	Status        AppointmentStatus `json:"status"`
	PaymentStatus PaymentStatus     `json:"payment_status"`
}

type AppointmentStatus string

const (
	AppointmentStatusConfirmed AppointmentStatus = "confirmed"
	AppointmentStatusPending   AppointmentStatus = "pending"
	AppointmentStatusCancelled AppointmentStatus = "cancelled"
	AppointmentStatusCompleted AppointmentStatus = "completed"
)

type PaymentStatus string

const (
	StatusPaid    PaymentStatus = "paid"
	StatusPending PaymentStatus = "pending"
	StatusFailed  PaymentStatus = "failed"
)

type TimeSlot struct {
	Id         string    `json:"id"`
	BusinessId string    `json:"business_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	IsBooked   bool      `json:"is_booked"`
}
