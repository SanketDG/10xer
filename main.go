package main

import "fmt"
import "math/rand"
import "time"
import "os"
import "gopkg.in/src-d/go-git.v4"
import "gopkg.in/src-d/go-git.v4/plumbing/object"

// TODO: After deploying on Kubernetes, we can OPTIMIZE to use a string
func return_go_code() string {
	return "fmt.Println('10X!!!!\n')"
}

func return_filename() string {
	return "10X.go"
}

func ONE_INDEXIFY(num int) int {
	return num+1
}

func get_random_number(num int) int {
	return ONE_INDEXIFY(rand.Intn(num))
}

func get_date_for_tenx_commit(one_year_ago_date time.Time) time.Time {
	// random_number  := random_number_one_to_five(5) FUTURE FEATURE, PAY IN BITCOIN FOR PRIVATE RELEASE HIHI
	date_for_tenx_commit := one_year_ago_date.AddDate(0, 0, 1)
	return date_for_tenx_commit
}

func make_date_format_ready() string {
	return "2010-10-10 10:10:10"
}

func get_one_year_in_numbers() int {
	return 365
}

func get_author(name string, email string, date_formatted string) *object.Signature {
	when, _ := time.Parse(object.DateFormat, date_formatted)
	return &object.Signature{
		Name:  name,
		Email: email,
		When:  when,
	}
}

func tenx_me_right_now(name string, email string) {
	// SINGLE RESPONSIBLITY PRINCIPLE, ONE FUNCTION ONE THING, NO ERROR CHECKS!
	tenx_countdowner := 0
	one_year_number_of_days := get_one_year_in_numbers()
	file_content := return_go_code()
	filename := return_filename()
	repository, _ := git.PlainInit("10x-project", false)
	worktree, _ := repository.Worktree()
	file, _ := os.Create("10x.go")
	now := time.Now()
	one_year_ago_date := now.AddDate(-1, 0, 0)
	date_for_tenx_commit := get_date_for_tenx_commit(one_year_ago_date)
WRTITE_TENX_CODE:
	tenx_countdowner += 1
	if tenx_countdowner == one_year_number_of_days { os.Exit(10) }
	file.WriteString(file_content)
	worktree.Add(filename)
	date := get_date_for_tenx_commit(date_for_tenx_commit)
	date_for_tenx_commit = date
	date_formatted := date.Format(object.DateFormat)
	fmt.Println(date_formatted)
	author := get_author(name, email, date_formatted)
	worktree.Commit("10X WOOHOOO!", &git.CommitOptions{Author: author})
	goto WRTITE_TENX_CODE


}

func main() {
	// I SHUD PARSE ARGS BUT FIRST NEED SOME 10X GITHUB STARS
    tenx_me_right_now("", "")
}