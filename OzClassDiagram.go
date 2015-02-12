package ozcld

import (
	"strings"
)

// Dot 形式の文字列に変換できることを表すインターフェース
type Dot interface {
	// Dot 形式の文字列を返却する
	ToDot() string
}

// フィールド
type Field struct {
	field string
}

// Dot 形式の文字列を返却する
func (this *Field) ToDot() string {
	return this.field + "\\l"
}

// フィールドを作成する
func CreateFieldFromString(def string) *Field {
	return &Field{def}
}

// フィールドリスト
type Fields []*Field

// Dot 形式の文字列を返却する
func (this Fields) ToDot() string {
	// 必要な長さのスライスを作成
	defs := make([]string, len(this))

	// スライスにフィールド定義を格納
	for i, v := range this {
		defs[i] = v.ToDot()
	}

	// Fields 返却
	return strings.Join(defs, "")
}

// メソッド
type Method struct {
	method string
}

// メソッドを作成する
func CreateMethodFromString(def string) *Method {
	return &Method{def}
}

// Dot 形式の文字列を返却する
func (this *Method) ToDot() string {
	return this.method + "\\l"
}

// メソッドリスト
type Methods []*Method

// Dot 形式の文字列を返却する
func (this Methods) ToDot() string {
	// 必要な長さのスライスを作成
	defs := make([]string, len(this))

	// スライスにフィールド定義を格納
	for i, v := range this {
		defs[i] = v.ToDot()
	}

	// Methods 返却
	return strings.Join(defs, "")
}

// クラス
type Class struct {
	id         string
	stereotype string
	name       string
	fields     Fields
	methods    Methods
}

// 各オブジェクトからクラスを作成する
func CreateClass(stereotype string, name string, fields Fields, methods Methods) *Class {
	if fields == nil {
		fields = Fields{}
	}
	if methods == nil {
		methods = Methods{}
	}
	return &Class{"", stereotype, name, fields, methods}
}

// 文字列からクラスを作成する
func CreateClassFromDefs(stereotype string, name string, fieldDefs []string, methodDefs []string) *Class {
	fields := createFieldsFromStrings(fieldDefs)
	methods := createMethodsFromStrings(methodDefs)
	return CreateClass(stereotype, name, fields, methods)
}

// 文字列のリストから Field リストを作成する
func createFieldsFromStrings(fieldDefs []string) Fields {
	fields := make(Fields, len(fieldDefs))

	for i, field := range fieldDefs {
		fields[i] = CreateFieldFromString(field)
	}

	return fields
}

// 文字列のリストから Method リストを作成する
func createMethodsFromStrings(methodDefs []string) Methods {
	methods := make(Methods, len(methodDefs))

	for i, method := range methodDefs {
		methods[i] = CreateMethodFromString(method)
	}

	return methods
}

// 識別文字列を設定する
func (this *Class) SetIdent(id string) {
	this.id = id
}

// 識別文字列を取得する
func (this *Class) GetIdent() string {
	if this.id != "" {
		return this.id
	}

	// 識別子に "." は使えないので "_" に置き換える
	return escapeChars(this.name)
}

// Dot 形式の文字列を返却する
func (this *Class) ToDot() string {
	// 必要な長さのスライスを作成
	defs := []string{this.GetIdent(), " [label = \"{"}

	if this.stereotype != "" {
		defs = append(defs, "\\<\\<", this.stereotype, "\\>\\>\\n")
	}

	defs = append(defs, this.name, "|", this.fields.ToDot(), "|", this.methods.ToDot(), "}\"];")

	// 文字列返却
	return strings.Join(defs, "")
}

// Field を追加
func (this *Class) AddFieldFromString(def string) {
	field := CreateFieldFromString(def)
	this.fields = append(this.fields, field)
}

// Method を追加
func (this *Class) AddMethodFromString(def string) {
	method := CreateMethodFromString(def)
	this.methods = append(this.methods, method)
}

// 名前空間(パッケージ)
type Namespace struct {
	name       string
	classes    []*Class
	namespaces []*Namespace
}

// Namespace を作成する
func CreateNamespace(name string, classes []*Class, namespaces []*Namespace) *Namespace {
	if classes == nil {
		classes = []*Class{}
	}
	if namespaces == nil {
		namespaces = []*Namespace{}
	}
	this := &Namespace{name, classes, namespaces}

	return this
}

// Dot 形式の文字列を返却する
func (this *Namespace) ToDot() string {
	defs := []string{"subgraph " + escapeChars("cluster_" + this.name) + " {"}

	defs = append(defs, "label = \""+this.name+"\";")

	for _, v := range this.namespaces {
		defs = append(defs, v.ToDot())
	}

	for _, v := range this.classes {
		defs = append(defs, v.ToDot())
	}

	defs = append(defs, "}")

	return strings.Join(defs, "\n")
}

// 識別文字列を取得する
func (this *Namespace) GetIdent() string {
	// 識別子に "." は使えないので "_" に置き換える
	return escapeChars(this.name)
}

// Class を追加
func (this *Namespace) AddClass(class *Class) {
	this.classes = append(this.classes, class)
}

// Namespace を追加
func (this *Namespace) AddNamespace(namespace *Namespace) {
	this.namespaces = append(this.namespaces, namespace)
}

type RelationType int

// 関係
type Relation struct {
	name             string
	relationType     RelationType
	fromClass        string
	toClass          string
	fromMultiplicity string
	toMultiplicity   string
}

// Relation を作成する
func CreateRelation(name string, relationType RelationType, fromClass string, toClass string, fromMultiplicity string, toMultiplicity string) *Relation {
	return &Relation{name, relationType, fromClass, toClass, fromMultiplicity, toMultiplicity}
}

// 関係の種類
const (
	RELATION_NORMAL RelationType = iota
	RELATION_INHERIT
	RELATION_IMPLEMENT
	RELATION_AGGREGATION
	RELATION_COMPOSITION
)

func (this *Relation) getEdgeStyles() (style string, arrowhead string) {

	if this.relationType == RELATION_INHERIT {
		style = "solid"
		arrowhead = "onormal"
	} else if this.relationType == RELATION_IMPLEMENT {
		style = "dashed"
		arrowhead = "onormal"
	} else if this.relationType == RELATION_AGGREGATION {
		style = "solid"
		arrowhead = ""
	} else if this.relationType == RELATION_COMPOSITION {
		style = "solid"
		arrowhead = ""
	}

	return style, arrowhead
}

// Dot 形式の文字列を返却する
func (this *Relation) ToDot() string {

	style, arrowhead := this.getEdgeStyles()

	// 基本
	base := []string{"edge [style = \"" + style + "\", arrowhead = \"" + arrowhead + "\"];\n"}
	base = append(base, this.fromClass+" -> "+this.toClass)

	// 詳細
	detail := []string{}

	// 関係名名前
	if this.name != "" {
		detail = append(detail, "label = \""+this.name+"\"")
	}

	// FROM の処理
	if this.fromMultiplicity != "" {
		detail = append(detail, "taillabel = \""+this.fromMultiplicity+"\"")
	}

	// TO の処理
	if this.toMultiplicity != "" {
		detail = append(detail, "headlabel = \""+this.toMultiplicity+"\"")
	}

	var detailString string
	if len(detail) != 0 {
		detailString = "[" + strings.Join(detail, ",") + "];"
	} else {
		detailString = ""
	}

	return strings.Join(base, "") + detailString
}

// クラス図
type ClassDiagram struct {
	name       string
	namespaces []*Namespace
	classes    []*Class
	relations  []*Relation
}

// クラス図作成
func CreateClassDiagram(name string, namespaces []*Namespace, classes []*Class, relations []*Relation) *ClassDiagram {
	return &ClassDiagram{name, namespaces, classes, relations}
}

// Dot 形式の文字列を返却する
func (this *ClassDiagram) ToDot() string {

	defs := []string{"digraph " + this.name + " {\nnode [shape = record];\n"}

	for _, v := range this.namespaces {
		defs = append(defs, v.ToDot())
	}

	for _, v := range this.classes {
		defs = append(defs, v.ToDot())
	}

	for _, v := range this.relations {
		defs = append(defs, v.ToDot())
	}

	defs = append(defs, "}")

	return strings.Join(defs, "\n")
}

// Class を追加
func (this *ClassDiagram) AddClass(class *Class) {
	this.classes = append(this.classes, class)
}

// Namespace を追加
func (this *ClassDiagram) AddNamespace(ns *Namespace) {
	this.namespaces = append(this.namespaces, ns)
}

// Relation を追加
func (this *ClassDiagram) AddRelation(relation *Relation) {
	this.relations = append(this.relations, relation)
}

func escapeChars(str string) string {
	return "\"" + str + "\""
}
