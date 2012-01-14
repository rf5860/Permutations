package permutations

import (
    "io/ioutil"
    "strings"
)

func ReadDictionary(file string) (words []string) {
    contents, _ := ioutil.ReadFile(file)
    words = strings.Split(string(contents), "\n")
    return words
}

func concat(first, second []([]byte)) (ret []([]byte)) {
   ret = make([]([]byte), len(first) + len(second))
   copy(ret, first)
   copy(ret[len(first):], second)
   return ret
}

func UpperStrings(words []string) (stringsUpper []string) {
    for _, word := range words {
        stringsUpper = append(stringsUpper, strings.ToUpper(word))
    }
    
    return
}

func Permute(data []byte, e int) (perms []([]byte)){
    if e <= 1 {
        r := make([]byte, len(data))
        copy(r, data)
        perms = append(perms, r)
        return
    }
    
    e1 := e - 1
    p := len(data) - e1
    perms = concat(perms, Permute(data, e1))
    dataw := make([]byte, len(data))
    copy(dataw, data)
    for i := p; i > 0; i-- {   
        dataw[i], dataw[i - 1] = dataw[i - 1], dataw[i]
        perms = concat(perms, Permute(dataw, e1))
    }
    
    return
}
// WARNING: Shouldn't be used with large files - V! permutations for N files means a very long run-time.
//  This func is a proof-of-concept only.
func MakeAnagrams(file string) (dict map[string][]([]byte)) {
    normWords := ReadDictionary(file)
    words := UpperStrings(normWords)
    dict = make(map[string][]([]byte))
    for _, word := range words {
        dict[word] = Permute([]byte(word), len(word))
    }
    
    for _, word := range words {
        ana := make([]([]byte), 1)
        for _, perm := range dict[word] {
            if dict[string(perm)] != nil {
                ana = append(ana, perm) 
            }
        }
        dict[word] = ana
    }
    
    return
}
