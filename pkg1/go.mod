module github.com/skyflow-bharti/sample-pkg/pkg1

go 1.20

replace github.com/skyflow-bharti/sample-pkg/v1/pkg2 latest => /Users/bharts/Desktop/sample-pkg/v1/pkg2

require github.com/skyflow-bharti/sample-pkg/v1/pkg2 latest
