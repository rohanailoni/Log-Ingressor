package Models

import "errors"

type Flagvalue struct {
	Level            Valueset
	Message          Valueset
	ResourceId       Valueset
	TraceId          Valueset
	SpanId           Valueset
	Commit           Valueset
	ParentResourceId Valueset
	Timestamp        string
	FromTimestamp    string
	ToTimestamp      string
}

func (this *Flagvalue) CheckDuplicateOnAllFlags() error {
	err := this.Level.checkDuplicate()
	if err != nil {
		return errors.New("duplicate query found in level field")
	}
	err = this.Message.checkDuplicate()
	if err != nil {
		return errors.New("duplicate query found in message field")
	}
	err = this.ResourceId.checkDuplicate()
	if err != nil {
		return errors.New("duplicate query found in resourceId field")
	}
	err = this.TraceId.checkDuplicate()
	if err != nil {
		return errors.New("duplicate query found in TraceId field")
	}
	err = this.SpanId.checkDuplicate()
	if err != nil {
		return errors.New("duplicate query found in spanId field")
	}
	err = this.Commit.checkDuplicate()
	if err != nil {
		return errors.New("duplicate query found in commit field")
	}
	err = this.ParentResourceId.checkDuplicate()
	if err != nil {
		return errors.New("duplicate query found in MetaData parentResourceId field")
	}

	if this.ToTimestamp != "" && this.FromTimestamp == "" {
		return errors.New("error fromtime not given but totime is given! try to give fromtime and totime automatically set to NOW")
	}
	if this.Timestamp != "" && (this.FromTimestamp != "" || this.ToTimestamp != "") {
		return errors.New("filter on both fromtime or totime and timeline is given this is not allowed")

	}
	return nil
}

// Valueset This structure defined for taking input of all type of flags
type Valueset struct {
	RegularFlag  string
	RegexFlag    string
	WildcardFlag string
}

func (this *Valueset) checkDuplicate() error {
	bitMask := 0
	if this.RegularFlag != "" {
		bitMask |= 1 << 0
	}
	if this.RegexFlag != "" {
		bitMask |= 1 << 1
	}
	if this.WildcardFlag != "" {
		bitMask |= 1 << 2
	}

	if bitMask == 1<<0 || bitMask == 1<<1 || bitMask == 1<<2 || bitMask == 0 {
		return nil
	} else {
		return errors.New("duplicate found")
	}

}

// HasValue if there is a value return true else false
func (this *Valueset) HasValue() bool {
	if this.RegularFlag == "" && this.WildcardFlag == "" && this.RegexFlag == "" {
		return false
	}
	return true
}
func (this *Valueset) HasRegex() bool {
	if this.RegexFlag == "" {
		return false
	}
	return true
}
func (this *Valueset) HasRegular() bool {
	if this.RegularFlag == "" {
		return false
	}
	return true
}
func (this *Valueset) HasWildcard() bool {
	if this.WildcardFlag == "" {
		return false
	}
	return true
}
