package proverb

func Proverb(rhyme []string) []string {
	newProverb := []string{}
	for i:=0; i<len(rhyme)-1; i++ {
		newProverb = append(newProverb, "For want of a "+rhyme[i]+" the "+rhyme[i+1]+" was lost.")
	}
	newProverb = append(newProverb, "And all for the want of a " + rhyme[0]+".")
	return newProverb
}
