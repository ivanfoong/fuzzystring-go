package fuzzystring

import (
  "fmt"
  "regexp"
  "github.com/ivanfoong/strikeamatch-go/strikeamatch"
  "github.com/ivanfoong/stringutils-go/stringutils"
)

func filterNonWord(words []string) (filtered []string) {
  for  _, word := range words {
    r, err := regexp.Compile(`\W+`)
    
    if err != nil {
      fmt.Printf("There is a problem with your regexp.\n")
      break
    }
    
    if r.FindStringSubmatch(word) == nil {
      filtered = append(filtered, word)
    }
  }
  return filtered
}

func filterMinChar(words []string, minCharCount int) (filtered []string) {
  for  _, word := range words {
    if len(word) >= minCharCount {
      filtered = append(filtered, word)
    }
  }
  return filtered
}

func scoreStrings(s1 []string, s2 []string) (score float64) {
  size1 := len(s1)
  size2 := len(s2)
  
  occurrence1 := stringutils.StringOccurrence(s1)
  occurrence2 := stringutils.StringOccurrence(s2)
  
  smallerOccurrence := occurrence2
  largerOccurrence := occurrence1
  if size1 < size2 {
    smallerOccurrence = occurrence1
    largerOccurrence = occurrence2
  }
  
  intersectionOccurrence := 0
  
  for s, smallerOccurrence := range smallerOccurrence {
    largerOccurrence, ok := largerOccurrence[s]
    if ok && largerOccurrence > 0 {
      if smallerOccurrence < largerOccurrence {
        intersectionOccurrence += smallerOccurrence
      } else {
        intersectionOccurrence += largerOccurrence
      }
    } 
  }
  
  return (2.0 * float64(intersectionOccurrence)) / float64(size1 + size2)
}

func CompareString(text1 string, text2 string) (score float64) {
  pairsScore := strikeamatch.CompareString(text1, text2)
  
  words1 := stringutils.Words(text1)
  words2 := stringutils.Words(text2)
  
  minCharCount := 2
  filteredWords1 := filterNonWord(filterMinChar(words1, minCharCount))
  filteredWwords2 := filterNonWord(filterMinChar(words2, minCharCount))
  
  wordsScore := scoreStrings(filteredWords1, filteredWwords2)
  
  return (0.8 * pairsScore) + (0.2 * wordsScore)
}