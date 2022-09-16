package main

import (
	"fmt"
	"io"
	"bufio"
	"os"
	"strings"
)

func ParseHeader(r io.Reader) (headers []string, seqs []string) {
	var curseq strings.Builder
	s := bufio.NewScanner(r)
	s.Buffer(make([]byte, 0), 1e12)
	for s.Scan() {
		if s.Text() == "" {
			continue
		}
		if s.Text()[0] == '>' {
			headers = append(headers, s.Text()[1:])
			if curseq.String() != "" {
				seqs = append(seqs, curseq.String())
			}
			curseq.Reset()
		} else {
			fmt.Fprintf(&curseq, "%s", s.Text())
		}
	}
	seqs = append(seqs, curseq.String())
	return
}

func WriteFaStruct(w io.Writer, headers []string, seqs []string) {
	fmt.Fprintf(w, `# Created by Octave 5.2.0, Wed Sep 07 15:06:00 2022 MDT <jgbaldwinbrown@jgbaldwinbrown-ThinkPad-T480>
`)
	fmt.Fprintf(
		w,
		`# name: Sequence
# type: matrix
# rows: 1
# columns: %v
`,
		len(seqs),
	)

	for _, seq := range seqs {
		fmt.Fprintf(w, "%d\n", len(seq))
	}

	fmt.Fprintf(w, "\n\n")
}

func FaToMl(r io.Reader, w io.Writer) {
	headers, seqs := ParseHeader(r)
	WriteFaStruct(w, headers, seqs)
}

func main() {
	bw := bufio.NewWriter(os.Stdout)
	defer bw.Flush()
	FaToMl(os.Stdin, bw)
}

/*
# Created by Octave 5.2.0, Tue Mar 29 17:40:12 2022 MDT <jgbaldwinbrown@jgbaldwinbrown-ThinkPad-T480>
# name: F
# type: scalar struct
# ndims: 2
 1 1
# length: 2
# name: Header
# type: sq_string
# elements: 2
# length: 20
1                   
# length: 20
two                 


# name: Sequence
# type: sq_string
# elements: 3
# length: 21
AGTCGTCAA            
# length: 21
GCACAGT              
# length: 21
ACGACGTCAGTCAGTCACGTA

*/
