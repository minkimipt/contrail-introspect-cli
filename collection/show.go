package collection

import "fmt"
import "log"
import "github.com/gosuri/uitable"
import "github.com/minkimipt/contrail-introspect-cli/utils"

// Several representations of resources
type Shower interface {
	Long(maxColWidth uint)
	Short()
	Xml()
	Vars(slice *[][]string)
}

func (e Element) Xml() {
	fmt.Printf("%s", e.Node)
}
func (elts Elements) Xml() {
	for _, e := range elts {
		e.Xml()
	}
}
func (c Collection) Xml() {
	fmt.Printf("%s", c.rootNode)
}

func (e Element) Short() {
	if e.desc.ShortDetailXpath == "" {
		log.Fatal("ShortDetailXpath is not defined. Use XML output instead (-x option)")
		return
	}
	s, _ := e.Node.Search(e.desc.ShortDetailXpath)
	if len(s) != 1 {
		return
		// TODO Debug mode
		//log.Fatal("Xpath '" + e.desc.ShortDetailXpath + "' is not valid (check your ShortDetailXpath definition)")
	}
	fmt.Printf("%s\n", s[0])
}
func (col Collection) Short() {
	Elements(col.elements).Short()
}
func (elts Elements) Short() {
	for _, e := range elts {
		e.Short()
	}
}
func (e Element) Long(maxColWidth uint) {
	table := uitable.New()
	table.MaxColWidth = maxColWidth
	e.desc.LongDetail.LongFormat(table, FORMAT_TEXT, e)
	fmt.Println(table)
}
func (e Element) Vars(slice *[][]string) {
}
func (col Collection) Long(maxColWidth uint) {
	switch col.descCol.DescElt.LongDetail.(type) {
	case LongFormatXpaths:
		vars := [][]string{}
		keys := col.descCol.DescElt.LongDetail.(LongFormatXpaths)
		Elements(col.elements).Vars(&vars)
		createTableFromVars(keys, vars)
	default:
		Elements(col.elements).Long(maxColWidth)
	}
}
func (col Collection) Vars(slice *[][]string) {
	switch col.descCol.DescElt.LongDetail.(type) {
	case LongFormatFn:
		log.Fatal("raw output is not provided for " + col.descCol.PageArgs[0])
		return
	default:
		Elements(col.elements).Vars(slice)
	}
}
func (elts Elements) Long(maxColWidth uint) {
	table := uitable.New()
	table.MaxColWidth = maxColWidth
	for i, e := range elts {
		format := FORMAT_TABLE
		if i == 0 {
			format = FORMAT_TABLE_HEADER
		}
		e.desc.LongDetail.LongFormat(table, format, e)
	}
	fmt.Println(table)
}
func (elts Elements) Vars(slice *[][]string) {
	var local_slice []string
	for _, e := range elts {
		local_slice = []string{}
		format := FORMAT_TEXT
		e.desc.LongDetail.LongFormat(&local_slice, format, e)
		*slice = append(*slice, local_slice)
	}
}

// This is used to show the long version of an Element.
type LongFormatter interface {
	LongFormat(t interface{}, f Format, e Element)
	GetFields() []string
}

type LongFormatFn (func(*uitable.Table, Element))
type LongFormatValuesFn (func(*[]string, Element))
type LongFormatXpaths []string

type Format uint8

const (
	FORMAT_TEXT         Format = 1
	FORMAT_TABLE_HEADER Format = 2
	FORMAT_TABLE        Format = 3
)

func (fn LongFormatFn) GetFields() []string {
	fmt.Println("This is not implemented")
	return []string{}
}

func (fn LongFormatValuesFn) GetFields() []string {
	fmt.Println("This is not implemented")
	return []string{}
}

func (xpaths LongFormatXpaths) GetFields() []string {
	return_slice := make([]string, len(xpaths))
	for i, xpath := range xpaths {
		return_slice[i] = xpath
	}
	return return_slice
}

func (fn LongFormatFn) LongFormat(table interface{}, format Format, e Element) {
	fn(table.(*uitable.Table), e)
}

func (fn LongFormatValuesFn) LongFormat(table interface{}, format Format, e Element) {
	fn(table.(*[]string), e)
}

func (xpaths LongFormatXpaths) LongFormat(table interface{}, format Format, e Element) {
	var local_slice *[]string
	local_slice = table.(*[]string)
	for _, xpath := range xpaths {
		s, _ := e.Node.Search(xpath + "/text()")
		*local_slice = append(*local_slice, (utils.Pretty(s)))
	}
}

func createTableFromVars(keys []string, values [][]string) {
	table := uitable.New()
	tmp := make([]interface{}, len(keys))
	for i, v := range keys {
		tmp[i] = v
	}
	table.AddRow(tmp...)
	for _, row := range values {
		tmp := make([]interface{}, len(row))
		for j, cell := range row {
			tmp[j] = cell
		}
		table.AddRow(tmp...)
	}
	fmt.Println(table)
}
func longFormatTable(table *uitable.Table, format Format, e Element, xpaths LongFormatXpaths) {
	if format == FORMAT_TABLE_HEADER {
		tmp := make([]interface{}, len(xpaths))
		for i, v := range xpaths {
			tmp[i] = v
		}
		table.AddRow(tmp...)
	}

	tmp := make([]interface{}, len(xpaths))
	for i, xpath := range xpaths {
		s, _ := e.Node.Search(xpath + "/text()")
		if len(s) == 1 {
			tmp[i] = utils.Pretty(s)
		}
	}
	table.AddRow(tmp...)
}
