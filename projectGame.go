package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Character interface {
	setAttack (int)
	setDefence (int)
	setHealth (int)
	Attack (character Character)
	firstSkill (character Character)
	ultimateSkill (character Character)
	Defence ()
	getDefence () int
	getHealth () int
	getAttack () int
	getFSChance () int
	getULTSChance () int
}

type Claymore struct {
	ATK int
	HP int
	DEF int
	fSkillChance int
	ultSkillChance int
	ID int
	Class string
}
type Catalyst struct {
	ATK int
	HP int
	DEF int
	fSkillChance int
	ultSkillChance int
	ID int
	Class string
}
type Archer struct {
	ATK int
	HP int
	DEF int
	fSkillChance int
	ultSkillChance int
	ID int
	Class string
}

/*CLAYMORE*/

func (claymore *Claymore) setHealth (HP int){
	claymore.HP = HP
}
func (claymore *Claymore) setDefence (DEF int){
	claymore.DEF = DEF
}
func (claymore *Claymore) setAttack (ATK int){
	claymore.ATK = ATK
}

func (claymore *Claymore) Attack (character Character){
	damage := claymore.ATK - character.getDefence()
	if damage > 0 {
		character.setHealth(character.getHealth() - damage)
	}
}
func (claymore *Claymore) firstSkill (character Character) {
	damage := claymore.ATK + 50%claymore.ATK - character.getDefence()
	if damage > 0 {
		character.setHealth(character.getHealth() - damage)
	}
	DEF := character.getDefence()
	if DEF > 0 {
		character.setDefence(DEF - 50%DEF)
	}
}
func (claymore *Claymore) ultimateSkill (character Character){
	damage := claymore.ATK + 150%claymore.ATK - character.getDefence()
	if damage > 0 {
		character.setHealth(character.getHealth() - damage)
	}
	if claymore.DEF > 0 {
		claymore.DEF -= 50 % claymore.DEF
	}
}

func (claymore *Claymore) Defence (){
	if claymore.DEF <= 0{
		claymore.DEF = 15
	} else {
		claymore.DEF += 25%claymore.DEF
	}
}// + 25%DEF

func (claymore *Claymore) getDefence () int {
	return claymore.DEF
}
func (claymore *Claymore) getHealth () int {
	return claymore.HP
}
func (claymore *Claymore) getAttack () int{
	return claymore.ATK
}
func (claymore *Claymore) getFSChance () int{
	return claymore.fSkillChance
}
func (claymore *Claymore) getULTSChance () int{
	return claymore.ultSkillChance
}

/*CATALYST*/

func (catalyst *Catalyst) setHealth (HP int){
	catalyst.HP = HP
}
func (catalyst *Catalyst) setDefence (DEF int){
	catalyst.DEF = DEF
}
func (catalyst *Catalyst) setAttack (ATK int){
	catalyst.ATK = ATK
}

func (catalyst *Catalyst) Attack (character Character){
	damage := catalyst.ATK - character.getDefence()
	if damage > 0 {
		character.setHealth(character.getHealth() - damage)
	}
}
func (catalyst *Catalyst) firstSkill (character Character){
	ATK := character.getAttack()
	if ATK > 0 {
		character.setAttack(ATK - 15%ATK)
	}
	damage := catalyst.ATK + 10%catalyst.ATK - character.getDefence()
	if damage > 0 {
		character.setHealth(character.getHealth() - damage)
	}
}
func (catalyst *Catalyst) ultimateSkill (character Character){
	damage := catalyst.ATK + 150%catalyst.ATK - character.getDefence()
	character.setHealth(character.getHealth() - damage)
	if catalyst.DEF > 0 {
		catalyst.DEF -= 15%catalyst.DEF
	}
}

func (catalyst *Catalyst) Defence (){
	if catalyst.DEF <=0 {
		catalyst.DEF = 15
	} else {
		catalyst.DEF += 25 % catalyst.DEF
	}
	if catalyst.ATK <= 0 {
		catalyst.ATK = 25
	} else {
		catalyst.ATK += 10 % catalyst.ATK
	}
}// + 25%DEF + 10%ATK

func (catalyst *Catalyst) getDefence () int {
	return catalyst.DEF
}
func (catalyst *Catalyst) getHealth () int {
	return catalyst.HP
}
func (catalyst *Catalyst) getAttack () int {
	return catalyst.ATK
}
func (catalyst *Catalyst) getFSChance () int{
	return catalyst.fSkillChance
}
func (catalyst *Catalyst) getULTSChance () int{
	return catalyst.ultSkillChance
}


/*ARCHER*/

func (archer *Archer) setHealth (HP int){
	archer.HP = HP
}
func (archer *Archer) setDefence (DEF int){
	archer.DEF = DEF
}
func (archer *Archer) setAttack (ATK int){
	archer.ATK = ATK
}

func (archer *Archer) Attack (character Character){
	damage := archer.ATK - character.getDefence()
	if damage > 0 {
		character.setHealth(character.getHealth() - damage)
	}
}
func (archer *Archer) firstSkill (character Character){
	archer.ATK += 15%archer.ATK
	damage := archer.ATK + 50%archer.ATK - character.getDefence()
	if damage > 0 {
		character.setHealth(character.getHealth() - damage)
	}

}
func (archer *Archer) ultimateSkill (character Character){
	damage := archer.ATK + 150%archer.ATK - character.getDefence()
	if damage > 0 {
		character.setHealth(character.getHealth() - damage)
	}
	DEF := character.getDefence()
	if DEF > 0 {
		character.setDefence(DEF - 25%DEF)
	}
	if archer.ATK > 0 {
		archer.ATK -= 20%archer.ATK
	}
}

func (archer *Archer) Defence (){
	if archer.DEF <= 0 {
		archer.DEF = 15
	} else {
		archer.DEF += 5%archer.DEF
	}
	if archer.ATK <= 0 {
		archer.ATK = 35
	} else {
		archer.ATK += 15 % archer.ATK
	}
} //+ 5%DEF + 15%ATK

func (archer *Archer) getDefence () int {
	return archer.DEF
}
func (archer *Archer) getHealth () int {
	return archer.HP
}
func (archer *Archer) getAttack () int {
	return archer.ATK
}
func (archer *Archer) getFSChance () int{
	return archer.fSkillChance
}
func (archer *Archer) getULTSChance () int{
	return archer.ultSkillChance
}


func Fight (c1 Character, c2 Character) Character {
	count := 0
	for c1.getHealth() > 0 && c2.getHealth() > 0 {
		if count == 10 {
			c1.Defence()
			c2.Defence()
			count = 0
		}
		n := 1 + rand.Intn(100)
		if n % c1.getFSChance() == 0 {
			c1.firstSkill(c2)
		}
		if n % c1.getULTSChance() == 0{
			c1.ultimateSkill(c2)
		} else {
			c1.Attack(c2)
		}

		n = 1 + rand.Intn(100)
		if n % c2.getFSChance() == 0 {
			c2.firstSkill(c1)
		}
		if n % c2.getULTSChance() == 0{
			c2.ultimateSkill(c1)
		} else {
			c2.Attack(c1)
		}
		count ++
	}
	if c1.getHealth() <= 0 {
		return c2
	}
	return c1
}

func FightResult (ch1 , ch2 chan Character, characters *[]Character){
	character1 := <- ch1
	character2 := <- ch2
	fmt.Println("Битва между: ", character1, " и ", character2)
	winner := Fight(character1, character2)
	fmt.Println("Победил: ", winner)
	*characters = append(*characters, winner)
}

func CharacterSelection (characters *[]Character, ch1, ch2 chan Character){
	if len(*characters) == 0 || len(*characters) == 1 {
		return
	}
	c1 := 0 + rand.Intn(len(*characters))
	ch1 <- (*characters)[c1]
	(*characters)[c1] = (*characters)[len(*characters) - 1]
	*characters = (*characters)[:len(*characters) - 1]
	c2 := 0 + rand.Intn(len(*characters))
	ch2 <- (*characters)[c2]
	(*characters)[c2] = (*characters)[len(*characters) - 1]
	*characters = (*characters)[:len(*characters) - 1]
}

func CreateCharactersArray (characters *[]Character, n int){
	for i := 0; i < n; i++ {
		(*characters)[i] = &Claymore{
			ATK:            100,
			HP:             1000,
			DEF:            50,
			fSkillChance:   5,
			ultSkillChance: 7,
			ID:             i,
			Class:			"claymore",
		}
		(*characters)[i + n] = &Archer{
			ATK:            70,
			HP:             1000,
			DEF:            50,
			fSkillChance:   3,
			ultSkillChance: 6,
			ID:             i + n,
			Class: 			"archer",
		}
		(*characters)[i + 2 * n] = &Catalyst{
			ATK:            60,
			HP:             1000,
			DEF:            50,
			fSkillChance:   2,
			ultSkillChance: 4,
			ID:             i + 2 * n,
			Class:			"catalyst",
		}
	}
}

func main() {
	var n int
	fmt.Println("Введите количество персонажей в одном классе: ")
	fmt.Scan(&n)
	characters := make([]Character, 3 * n)
	CreateCharactersArray(&characters, n)

	ch1 := make (chan Character)
	ch2 := make (chan Character)

	for len(characters) != 1{
		go CharacterSelection(&characters, ch1, ch2)
		go FightResult(ch1, ch2, &characters)
		time.Sleep(3)
	}

	fmt.Println("В живых остался: ", characters[0])
}