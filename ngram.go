package ngram

const (
    minNgram int = 3
)

func MakeNgram(name string) []string{
    ngram := make([]string, 0)
    if len(name) < minNgram {
        return append(ngram, name)
    }
    for i:= minNgram; i <= len(name); i++ {
        for j := 0; j <= len(name) - i; j++ {
            ngram = append(ngram, name[j:j + i])
        }
    }

    return ngram
}

