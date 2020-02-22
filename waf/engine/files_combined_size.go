package engine

import (
	"github.com/asalih/guardian/matches"
)

var FILES_COMBINED_SIZE = "FILES_COMBINED_SIZE"

func (t *TransactionMap) loadFilesCombinedSize() *TransactionMap {
	t.variableMap[FILES_COMBINED_SIZE] =
		&TransactionData{func(executer *TransactionExecuterModel) *matches.MatchResult {
			matchResult := matches.NewMatchResult()

			muliErr := executer.transaction.Request.ParseMultipartForm(1024 * 1024 * 4)

			if muliErr != nil {
				return matchResult.SetMatch(true)
			}

			files := executer.transaction.Request.MultipartForm.File
			totalSize := int64(0)
			for _, headers := range files {
				for _, head := range headers {
					totalSize += head.Size
				}
			}

			return executer.rule.ExecuteRule(totalSize)

		}}

	return t
}