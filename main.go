package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/hako/durafmt"
	"github.com/rivo/tview"
	"time"
)

type Activity struct {
	description string
	duration    time.Duration
}

var activities = make([]Activity, 0)

// Tview
var pages = tview.NewPages()
var activityText = tview.NewTextView()
var app = tview.NewApplication()
var activityForm = tview.NewForm()
var durationForm = tview.NewForm()
var activitiesList = tview.NewList().ShowSecondaryText(true)
var flex = tview.NewFlex()
var menuText = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText("(a) add activity - (q) quit")

var selected *Activity

func main() {
	activitiesList.SetSelectedFunc(func(index int, name string, secondName string, shortcut rune) {
		selected = &activities[index]
		durationForm.Clear(true)
		addDurationForm()
		pages.SwitchToPage("Add Duration")
	})

	flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().
			AddItem(activitiesList, 0, 1, true).
			AddItem(activityText, 0, 4, false), 0, 6, false).
		AddItem(menuText, 0, 1, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		} else if event.Rune() == 97 {
			activityForm.Clear(true)
			addActivityForm()
			pages.SwitchToPage("Add Activity")
		}
		return event
	})

	pages.AddPage("Menu", flex, true, true)
	pages.AddPage("Add Activity", activityForm, true, false)
	pages.AddPage("Add Duration", durationForm, true, false)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func addActivityList() {
	activitiesList.Clear()
	for index, activity := range activities {
		var dur = durafmt.Parse(activity.duration)
		activitiesList.AddItem(activity.description, dur.String(), rune(49+index), nil)
	}
}

func addActivityForm() *tview.Form {

	activity := Activity{}

	activityForm.AddInputField("what did you do?", "", 20, nil, func(description string) {
		activity.description = description
	})

	activityForm.AddButton("Save", func() {
		activities = append(activities, activity)
		addActivityList()
		pages.SwitchToPage("Menu")
	})

	return activityForm
}

func addDurationForm() *tview.Form {

	durationForm.AddInputField("edit duration?", "", 20, nil, func(duration string) {
		// TODO handle parsing error
		selected.duration, _ = time.ParseDuration(duration)
	})

	durationForm.AddButton("Save", func() {
		addActivityList()
		pages.SwitchToPage("Menu")
	})

	return durationForm
}
