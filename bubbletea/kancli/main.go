package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status int


const divisor = 4

const (
todo status = iota 
inProgress
done
)

/* CUSTOM ITEM */

type Task struct {
	status status
	title string
	description string
	loaded bool
}

//implement the list.Item interface

func(t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

/*MAIN MODEL*/

type Model struct {
	loaded bool
	focused status
	lists []list.Model
	err error
}

func New() *Model {
	return &Model{}
}


// TODO: call this on tea.WindowSizeMsg
func (m *Model) initLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/divisor-2, height)
	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList, defaultList}
	//Init ToDo
	m.lists[todo].Title = "To Do"
	m.lists[todo].SetItems([]list.Item{
		Task{status: todo, title: "buy milk", description: "strawberry milk"},
		Task{status: todo, title: "eat sushi", description: "negitoro roll, miso soup, rice"},
		Task{status: todo, title: "fold laundry", description: "or wear wrinkly t-shirts"},
		
	})


//Init in progress
m.lists[inProgress].Title = "In Progress"
m.lists[inProgress].SetItems([]list.Item{
	Task{status: todo, title: "write code", description: "don`t worry, its Go"},

	
})

//Init done
m.lists[done].Title = "Done"
m.lists[done].SetItems([]list.Item{
	Task{status: todo, title: "stay cool", description: "as a cucumber"},

	
})



}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg:= msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {

			m.initLists(msg.Width, msg.Height)
			m.loaded = true
		}
		
	}
	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.loaded {
		return lipgloss.JoinHorizontal(
			lipgloss.Left,
			m.lists[todo].View(),
			m.lists[inProgress].View(),
			m.lists[done].View(),
		)
	} else {
		return "loading..."
	}
	
}

func main() {
m := New()
p := tea.NewProgram(m)
if _,err := p.Run(); err != nil {
	fmt.Println(err)
	os.Exit(1)
}
}

