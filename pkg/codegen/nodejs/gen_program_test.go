package nodejs

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/codegen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/testing/test"
)

// For better CI test to job distribution, we split the test cases into three tests.

var genProgramBatchSize = len(test.PulumiPulumiProgramTests) / 3

func TestGenerateProgramOne(t *testing.T) {
	t.Parallel()

	testGenerateProgramBatch(t, test.PulumiPulumiProgramTests[0:genProgramBatchSize])
}

func TestGenerateProgramTwo(t *testing.T) {
	t.Parallel()

	testGenerateProgramBatch(t, test.PulumiPulumiProgramTests[genProgramBatchSize:2*genProgramBatchSize])
}

func TestGenerateProgramThree(t *testing.T) {
	t.Parallel()

	testGenerateProgramBatch(t, test.PulumiPulumiProgramTests[2*genProgramBatchSize:])
}

func testGenerateProgramBatch(t *testing.T, testCases []test.ProgramTest) {
	test.TestProgramCodegen(t,
		test.ProgramCodegenOptions{
			Language:   "nodejs",
			Extension:  "ts",
			OutputFile: "index.ts",
			Check: func(t *testing.T, path string, dependencies codegen.StringSet) {
				Check(t, path, dependencies, true)
			},
			GenProgram: GenerateProgram,
			TestCases:  testCases,
		})
}

func TestGenerateProgramVersionSelection(t *testing.T) {
	t.Parallel()

	expectedVersion := map[string]test.PkgVersionInfo{
		"aws-resource-options": {
			Pkg:          "\"@pulumi/aws\"",
			OpAndVersion: "\"4.38.0\"",
		},
	}

	test.TestProgramCodegen(t,
		test.ProgramCodegenOptions{
			Language:   "nodejs",
			Extension:  "ts",
			OutputFile: "index.ts",
			Check: func(t *testing.T, path string, dependencies codegen.StringSet) {
				Check(t, path, dependencies, true)
			},
			GenProgram: GenerateProgram,
			TestCases: []test.ProgramTest{
				{
					Directory:   "aws-resource-options",
					Description: "Resource Options",
					MockPluginVersions: map[string]string{
						"aws": "4.38.0",
					},
				},
			},

			IsGenProject:    true,
			GenProject:      GenerateProject,
			ExpectedVersion: expectedVersion,
			DependencyFile:  "package.json",
		})
}
