package day4

import (
	"fmt"
	"log"
	"testing"
)

func TestFindAllPasswords(t *testing.T) {
	t.Run("check FindAllPasswords", func(t *testing.T) {
		got := FindAllPasswords("111111", "111112")
		want := 2
		if got != want {
			log.Fatalf("want: %d, got: %d", want, got)
		}
	})

	t.Run("check hasAdjacentNumber", func(t *testing.T) {
		got := hasAdjacentNumber(111111)
		want := true
		if got != want {
			log.Fatalf("want: %t, got: %t", want, got)
		}
	})

	t.Run("get part 1 result", func(t *testing.T) {
		got := FindAllPasswords("197487", "673251")
		fmt.Println(got)
	})
}

func TestFindMoreStrictAllPasswords(t *testing.T) {
	t.Run("check FindMoreStrictAllPasswords", func(t *testing.T) {
		got := notLargerGroupMatching(112233)
		want := true
		if got != want {
			log.Fatalf("want1: %t, got: %t", want, got)
		}

		got = notLargerGroupMatching(123444)
		want = false
		if got != want {
			log.Fatalf("want2: %t, got: %t", want, got)
		}

		got = notLargerGroupMatching(111122)
		want = true
		if got != want {
			log.Fatalf("want3: %t, got: %t", want, got)
		}

		got = notLargerGroupMatching(222344)
		want = true
		if got != want {
			log.Fatalf("want4: %t, got: %t", want, got)
		}

		got = notLargerGroupMatching(225556)
		want = false
		if got != want {
			log.Fatalf("want5: %t, got: %t", want, got)
		}
	})

	t.Run("get part 2 result", func(t *testing.T) {
		got := FindMoreStrictAllPasswords("197487", "673251")
		fmt.Println(got)
	})
}
