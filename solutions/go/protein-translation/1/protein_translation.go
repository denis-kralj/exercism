package protein

import "errors"

var ErrStop = errors.New("")
var ErrInvalidBase = errors.New("")

func FromRNA(rna string) ([]string, error) {
    if len(rna) % 3 != 0 {
        return nil, ErrInvalidBase
    }

    result := []string{}

    for i:=0; i < len(rna); i+=3 {
        protein, err := FromCodon(rna[i:i+3])

        if err != nil{
			if err == ErrStop {
                return result, nil
            }
        	return nil, err
        }

        result = append(result, protein)
    }

    return result, nil
}

func FromCodon(codon string) (string, error) {
	switch codon {
        case "AUG": return "Methionine", nil
        case "UUU", "UUC": return "Phenylalanine", nil
    	case "UUA", "UUG": return "Leucine", nil
    	case "UCU", "UCC", "UCA", "UCG": return "Serine", nil
    	case "UAU", "UAC": return "Tyrosine", nil
    	case "UGU", "UGC": return "Cysteine", nil
    	case "UGG": return "Tryptophan", nil
    	case "UAA", "UAG", "UGA": return "", ErrStop
    	default: return "", ErrInvalidBase
    }
}
