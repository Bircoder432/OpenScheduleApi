package dto

import "time"

type Lesson struct {
	LessonID       uint      `json:"lessonId"`
	Title          string    `json:"title"`
	Cabinet        string    `json:"cabinet"`
	Date           time.Time `json:"date"`
	Teacher        string    `json:"teacher"`
	Order          uint      `json:"order"`
	StudentGroupID uint      `json:"studentGroupID"`
}

type ScheduleLessonResponse struct {
	Title     string     `json:"title"`
	Cabinet   string     `json:"cabinet"`
	Teacher   string     `json:"teacher"`
	Order     uint       `json:"order"`
	StartTime HourMinute `json:"startTime"`
	EndTime   HourMinute `json:"endTime"`
}

type ScheduleResponse struct {
	GroupID uint                     `json:"groupId"`
	Date    time.Time                `json:"date"`
	Lessons []ScheduleLessonResponse `json:"lessons"`
}
