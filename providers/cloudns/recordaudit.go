package cloudns

import (
	"github.com/StackExchange/dnscontrol/v3/models"
	"github.com/StackExchange/dnscontrol/v3/pkg/recordaudit"
)

// AuditRecordSupport returns an error if any records are not
// supportable by this provider.
func AuditRecordSupport(records []*models.RecordConfig) error {
	var err error

	err = recordaudit.TxtBackticks(records)
	if err != nil {
		return err
	}

	err = recordaudit.TxtEmpty(records)
	if err != nil {
		return err
	}

	err = recordaudit.TxtTrailingSpace(records)
	if err != nil {
		return err
	}

	return nil
}
