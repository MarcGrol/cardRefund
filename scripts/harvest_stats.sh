#!/bin/bash


if [ -z "${DUXXIEPLATFORM_HOME}" ]
then
    DUXXIEPLATFORM_HOME="${GOPATH}/src/github.com/Duxxie/platform"
fi

TOTAL_LINE_COUNT=$(find ${DUXXIEPLATFORM_HOME}  -name "*.go" -exec cat {} \; |wc -l)
TEST_REPORT_LINE_COUNT=$(find ${DUXXIEPLATFORM_HOME}  -name "testResults.go" -exec cat {} \; |wc -l)
HTML_LINE_COUNT=$(find ${DUXXIEPLATFORM_HOME}  -name "*_html.go" -exec cat {} \; |wc -l)
TOTAL_EFFECTIVE_LINE_COUNT_WITH_REPORTS_AND_HTML=$(find ${DUXXIEPLATFORM_HOME}  -name "*.go" -exec cat {} \; |wc -l)

TOTAL_EFFECTIVE_LINE_COUNT=$(expr ${TOTAL_EFFECTIVE_LINE_COUNT_WITH_REPORTS_AND_HTML} - ${TEST_REPORT_LINE_COUNT} - ${HTML_LINE_COUNT})

GENERATED_LINE_COUNT_WITH_REPORTS_AND_HTML=$(find ${DUXXIEPLATFORM_HOME} -name "*.go"  -exec grep -l '// Generated automatically' {} \; | xargs cat |wc -l)
GENERATED_LINE_COUNT=$(expr ${GENERATED_LINE_COUNT_WITH_REPORTS_AND_HTML} - ${TEST_REPORT_LINE_COUNT} - ${HTML_LINE_COUNT})
TOTAL_MANUAL_LINE_COUNT=$(expr ${TOTAL_EFFECTIVE_LINE_COUNT} - ${GENERATED_LINE_COUNT})

TOTAL_TEST_LINE_COUNT=$(find ${DUXXIEPLATFORM_HOME}  -name "*_test.go" -exec cat {} \; |wc -l)
GENERATED_TEST_LINE_COUNT=$(find ${DUXXIEPLATFORM_HOME} -name "*_test.go"  -exec grep -l '// Generated automatically' {} \; | xargs cat |wc -l)
MANUAL_TEST_LINE_COUNT=$(expr ${TOTAL_TEST_LINE_COUNT} - ${GENERATED_TEST_LINE_COUNT})

MANUAL_PROD_LINE_COUNT=$(expr ${TOTAL_MANUAL_LINE_COUNT} - ${MANUAL_TEST_LINE_COUNT})

REST_SERVICE_COUNT=$(find ${DUXXIEPLATFORM_HOME} -name "*.go" -exec grep '@RestService' {} \; |wc -l)
REST_SERVICE_OPERATION_COUNT=$(find ${DUXXIEPLATFORM_HOME} -name "*.go" -exec grep '@RestOperation' {} \; |wc -l)
EVENT_COUNT=$(find ${DUXXIEPLATFORM_HOME} -name "*.go" -exec grep '@Event(' {} \; |wc -l)
EVENT_HANDLER_COUNT=$(find ${DUXXIEPLATFORM_HOME} -name "*.go" -exec grep '@EventOperation(' {} \; |wc -l)

echo "Statistics on 'duxxie-platform'-backend"
echo "======================================="
printf "%30s: %d\n" "TOTAL_LINE_COUNT" "${TOTAL_LINE_COUNT}" 
echo ""
printf "%30s: %d (%s)\n" "TOTAL_EFFECTIVE_LINE_COUNT" "${TOTAL_EFFECTIVE_LINE_COUNT}" "without html and test-reports"
printf "%30s: %d\n" "GENERATED_LINE_COUNT" "${GENERATED_LINE_COUNT}"
printf "%30s: %d\n" "TOTAL_MANUAL_LINE_COUNT" "${TOTAL_MANUAL_LINE_COUNT}"
echo ""
printf "%30s: %d\n" "MANUAL_PROD_LINE_COUNT" "${MANUAL_PROD_LINE_COUNT}"
echo ""
printf "%30s: %d\n" "TOTAL_TEST_LINE_COUNT" "${TOTAL_TEST_LINE_COUNT}"
printf "%30s: %d\n" "GENERATED_TEST_LINE_COUNT" "${GENERATED_TEST_LINE_COUNT}"
printf "%30s: %d\n" "MANUAL_TEST_LINE_COUNT" "${MANUAL_TEST_LINE_COUNT}"
echo ""
printf "%30s: %d\n" "REST_SERVICE_COUNT" "${REST_SERVICE_COUNT}"
printf "%30s: %d\n" "REST_SERVICE_OPERATION_COUNT" "${REST_SERVICE_OPERATION_COUNT}"
printf "%30s: %d\n" "EVENT_COUNT" "${EVENT_COUNT}"
printf "%30s: %d\n" "EVENT_HANDLER_COUNT" "${EVENT_HANDLER_COUNT}"
echo ""
printf "%30s: %d\n" "TEST_REPORT_LINE_COUNT" "${TEST_REPORT_LINE_COUNT}"
printf "%30s: %d\n" "HTML_LINE_COUNT" "${HTML_LINE_COUNT}"
