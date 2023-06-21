package strategy

import (
	"fmt"
	"strings"
)

/**
Many algorithms can be decomposed into higher - and lower - level parts
Making tea can be decomposed into
	- The process of making a hot beverage
	- Tea-specific things
The high level algorithm can then be reused for making coffee or hot chocolate

Summary
- Seperates an algorithm into its 'skeleton' and concrete implementation steps,
which can be varied at run-time.
*/

type OutputFormat int

const (
	MarkDown OutputFormat = iota
	Html
)

type ListStrategy interface {
	Start(builder *strings.Builder)
	End(builder *strings.Builder)
	AddListItem(builder *strings.Builder, item string)
}

type MarkdownListStrategy struct{}

func (m *MarkdownListStrategy) Start(builder *strings.Builder) {
}

func (m *MarkdownListStrategy) End(builder *strings.Builder) {
}

func (m MarkdownListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString(" * " + item + "\n")
}

var _ ListStrategy = (*MarkdownListStrategy)(nil)

type HtmlListStrategy struct{}

func (h *HtmlListStrategy) Start(builder *strings.Builder) {
	builder.WriteString("<ul>\n")
}

func (h *HtmlListStrategy) End(builder *strings.Builder) {
	builder.WriteString("</ul>\n")
}

func (h *HtmlListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString("\t <li>" + item + "</li>\n")
}

var _ ListStrategy = (*HtmlListStrategy)(nil)

type TextProcessor struct {
	builder strings.Builder
	list    ListStrategy
}

func NewTextProcessor(list ListStrategy) *TextProcessor {
	return &TextProcessor{builder: strings.Builder{}, list: list}
}

func (t *TextProcessor) SetOutputFormat(fmt OutputFormat) {
	switch fmt {
	case MarkDown:
		t.list = &MarkdownListStrategy{}
	case Html:
		t.list = &HtmlListStrategy{}
	}
}

func (t *TextProcessor) AppendList(items []string) {
	s := t.list
	s.Start(&t.builder)

	for _, item := range items {
		s.AddListItem(&t.builder, item)
	}

	s.End(&t.builder)
}

func (t *TextProcessor) Reset() {
	t.builder.Reset()
}

func (t *TextProcessor) String() string {
	return t.builder.String()
}

/**
- Define an algorithm at a high level
- Define the interface you expect each strategy to follow
- Support the injection of the strategy into the high-level algorithm
*/

func Start() {
	t := NewTextProcessor(&MarkdownListStrategy{})
	t.AppendList([]string{"park", "gui", "woo"})

	fmt.Println(t)

	t.Reset()

	t.SetOutputFormat(Html)
	t.AppendList([]string{"park", "gui", "woo"})
	fmt.Println(t)
}
