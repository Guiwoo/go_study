package chain

import "fmt"

/**
A chain of components who all get a chance to process a command or a query,
optionally having default processing implementation and an ability to terminate the processing chain.
*/

type Creature struct {
	Name            string
	Attack, Defense int
}

func NewCreature(name string, attack, defense int) *Creature {
	return &Creature{Name: name, Attack: attack, Defense: defense}
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack, c.Defense)
}

type Modifier interface {
	Add(m Modifier)
	Handle()
}

type CreatureModifier struct {
	creature *Creature
	next     Modifier
}

func NewCreatureModifier(creature *Creature) *CreatureModifier {
	return &CreatureModifier{creature: creature}
}

func (c *CreatureModifier) Add(m Modifier) {
	if c.next != nil {
		c.next.Add(m)
	} else {
		c.next = m
	}
}

func (c CreatureModifier) Handle() {
	if c.next != nil {
		c.next.Handle()
	}
}

type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(c *Creature) *DoubleAttackModifier {
	return &DoubleAttackModifier{CreatureModifier{creature: c}}
}

func (d *DoubleAttackModifier) Handle() {
	fmt.Println("Doubling", d.creature.Name, "\b's attack")
	d.creature.Attack *= 2
	d.CreatureModifier.Handle()
}

type IncreaseDefenseModifier struct {
	CreatureModifier
}

func NewIncreaseDefenseModifier(c *Creature) *IncreaseDefenseModifier {
	return &IncreaseDefenseModifier{CreatureModifier{creature: c}}
}
func (i *IncreaseDefenseModifier) Handle() {
	fmt.Println("Increasing", i.creature.Name, "\b's defense")
	i.creature.Defense *= 3
	i.CreatureModifier.Handle()
}

type NoBonusesModifier struct {
	CreatureModifier
}

func NewNoBonusesModifier(c *Creature) *NoBonusesModifier {
	return &NoBonusesModifier{CreatureModifier{creature: c}}
}
func (n *NoBonusesModifier) Handle() {
	fmt.Println("No Bonuses for you!")
}

func Start() {
	goblin := NewCreature("Goblin", 1, 1)
	fmt.Println(goblin.String())

	root := NewCreatureModifier(goblin)

	root.Add(NewNoBonusesModifier(goblin))

	root.Add(NewDoubleAttackModifier(goblin))
	root.Add(NewDoubleAttackModifier(goblin))
	root.Add(NewIncreaseDefenseModifier(goblin))

	root.Handle()
	fmt.Println(goblin.String())
}
