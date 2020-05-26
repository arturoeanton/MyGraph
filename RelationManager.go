package main

var (
	graph map[string]RelationManager
)


// Operation is
type Operation func(e1, e2 MemSet) MemSet

type relationsType struct {
	relations map[string]MemSet
}

// RelationManager is
type RelationManager struct {
	objects map[string]relationsType
	ids     map[string]uint
	names   map[uint]string
}

// NewRelationManager is contructor
func NewRelationManager() RelationManager {
	return RelationManager{
		objects: make(map[string]relationsType),
		ids:     make(map[string]uint),
		names:   make(map[uint]string),
	}
}

// RemoveObject is
func (r RelationManager) RemoveObject(element string) bool {
	var id uint
	if _, ok := r.ids[element]; !ok {
		return true
	}
	id = r.ids[element]
	for tag, set := range r.objects[element].relations {
		for _, value := range set.Array() {
			var name string
			name = r.names[value]
			r.objects[name].relations[tag].Remove(id)
			r.objects[element].relations[tag].Remove(value)
		}
	}
	//delete(r.objects, element)
	//delete(r.names, id)
	//delete(r.ids, element)

	return true
}

// AddObject is
func (r RelationManager) AddObject(element string) bool {
	var id uint
	if _, ok := r.ids[element]; ok {
		return false
	}
	id = uint(len(r.ids))
	r.ids[element] = id
	r.names[id] = element
	r.objects[element] = relationsType{relations: make(map[string]MemSet)}
	return true
}

// AddTag is
func (r RelationManager) AddTag(element1, tag string) bool {
	r.AddObject(element1)
	if _, ok := r.objects[element1].relations[tag]; ok {
		return true
	}
	r.objects[element1].relations[tag] = NewMemSet()
	return true
}

// Relation is
func (r RelationManager) RelationBi(element1, tag, element2 string) bool {
	r.RelationR(element1,tag,element2)
	r.RelationL(element1,tag,element2)
	return true
}

// Relation is
func (r RelationManager) RelationR(element1, tag, element2 string) bool {
	tagReverse := "_INTERNAL@[REVERSE-HIDE]-"+tag
	r.AddObject(element1)
	r.AddObject(element2)
	r.AddTag(element1, tag)
	r.AddTag(element2, tagReverse)
	r.objects[element1].relations[tag].Add(r.ids[element2])
	r.objects[element2].relations[tagReverse].Add(r.ids[element1])
	return true
}

// Relation is
func (r RelationManager) RelationL(element1, tag, element2 string) bool {
	tagReverse := "_INTERNAL@[REVERSE-HIDE]-"+tag
	r.AddObject(element1)
	r.AddObject(element2)
	r.AddTag(element2, tag)
	r.AddTag(element1, tagReverse)
	r.objects[element2].relations[tag].Add(r.ids[element1])
	r.objects[element1].relations[tagReverse].Add(r.ids[element2])
	return true
}

// Relation is
func (r RelationManager) RemoveRelation(element1, tag, element2 string) bool {
	tagReverse := "_INTERNAL@[REVERSE-HIDE]-"+tag
	ms1,ok := r.objects[element1].relations[tag]
	if ok {
		ms1.Remove(r.ids[element2])
	}
	ms2, ok  :=r.objects[element2].relations[tag]
	if ok {
		ms2.Remove(r.ids[element1])
	}
	msr1,ok := r.objects[element1].relations[tagReverse]
	if ok {
		msr1.Remove(r.ids[element2])
	}
	msr2, ok  :=r.objects[element2].relations[tagReverse]
	if ok {
		msr2.Remove(r.ids[element1])
	}
	return true
}

// IntersectionPivotFirst is
func (r RelationManager) IntersectionPivotFirst(param ...string) MemSet {
	var tag, name string
	var e1, e2, e3 MemSet
	tag = param[0]
	name = param[1]
	e1 = r.objects[name].relations[tag]
	for index := 2; index < len(param); index++ {
		name = param[index]
		e2 = r.objects[name].relations[tag]
		e3 = e3.Union(e1.Intersection(e2))
	}
	return e3
}

// Evaluate is
func (r RelationManager) Evaluate(f Operation, param []string) MemSet {
	var tag, name string
	var e1, e2 MemSet
	tag = param[0]
	name = param[1]
	e1 = r.objects[name].relations[tag]

	for index := 2; index < len(param); index++ {
		name = param[index]
		e2 = r.objects[name].relations[tag]
		e1 = f(e1, e2)
	}
	return e1
}

// Intersection is
func (r RelationManager) Intersection(param ...string) MemSet {
	operation := func(e1, e2 MemSet) MemSet { return e1.Intersection(e2) }
	return r.Evaluate(operation, param)
}

// Union is
func (r RelationManager) Union(param ...string) MemSet {
	operation := func(e1, e2 MemSet) MemSet { return e1.Union(e2) }
	return r.Evaluate(operation, param)
}

// Diff is
func (r RelationManager) Diff(param ...string) MemSet {
	operation := func(e1, e2 MemSet) MemSet { return e1.Diff(e2) }
	return r.Evaluate(operation, param)
}

// SymDiff is
func (r RelationManager) SymDiff(param ...string) MemSet {
	operation := func(e1, e2 MemSet) MemSet {
		return e1.SymDiff(e2)
	}
	return r.Evaluate(operation, param)
}

// GetElementAll is
func (r RelationManager) GetElementAll(e1 MemSet) []string {
	var result []string
	for _, val := range e1.Array() {
		result = append(result, r.names[val])
	}
	return result
}