package sophos

const (
	// XRestdErrAck is an http.Header key.
	// As described, there are multiple RESTful API interactions that can fail due to incon
	// sistencies, e.g., object A references object B but object B is deleted. The RESTful API
	// will prevent damage and inconsistency to confd by returning an error.
	//
	// You can resolve this problem by removing the reference to object B. You can configure
	// confd to Do this automatically by setting the header value to “last”.
	// X-Restd-Err-Ack: last
	// To enable this globally for all non-fatal errors, you can set the header value to “all.”
	// X-Restd-Err-Ack: all
	//
	// Note – Use this setting only when you’re not deleting important data or objects. Other
	// wise, you can acknowledge the error and cancel the operation by setting the header
	// value to “none”. You can then troubleshoot the error with all data and objects saved.
	// X-Restd-Err-Ack: none
	//
	// Refer to Section 3.4 of Sophos documentation
	XRestdErrAck = "X-Restd-Err-Ack"

	// XRestdLockOverride is an http.Header key.
	// The confd object model supports locking objects to avoid unintended changes. To
	// check if an object is locked or unlocked, you can use the GET method. The response
	// will indicate the specific lock level, i.e., “_locked” can be set to “global”, “user”, or “”
	// (empty string). A “global” value is a system lock and cannot be changed. A “user” value
	// is a lock set by a specific user while an empty string indicates that no lock is set.
	//
	// In order to change “user” lock values, you can change the lock value to an empty string
	// and then modify the “user” value. This procedure involves three API calls and can be
	// error prone if these calls are interrupted or not completed in the correct order. The
	// lock override header allows the users to manually override the current value without
	// having to change the value multiple times.
	//
	// X-Restd-Lock-Override: yes
	//
	// Refer to Section 3.5 of Sophos documentation
	XRestdLockOverride = "X-Restd-Lock-Override"

	// XRestdInsert is an http.Header key.
	// In many cases when you create a new object, the object needs to be directly inserted
	// into a node in order to be active. This usually takes two operations; however, there is
	// an additional header you can use when creating objects to automatically activate the
	// objects. The header will insert a reference at the given position inside the node.
	//
	// For example, if you create a “packetfilter/packetfilter” rule object and need to add the
	// object into the “packetfilter.rules” node and make it active, you can use the “1” flag to
	// set the rule as the first rule. If you need to add the new rule object as the last rule, you
	// can use “-1” flag. If you need the new rule object to reside somewhere in the middle,
	// e.g., fourth rules, you can set the flag to “4”.
	//
	// X-Restd-Insert: packetfilter.rules 4
	//
	// Refer to Section 3.6 of Sophos documentation
	XRestdInsert = "X-Restd-Insert"

	// XRestdSession is an http.Header key.
	// Each interaction with the confd creates or reuses a session. Sessions are important
	// for validation interaction and performance. However, maintaining sessions are
	// resource intensive and can degrade performance. If you use the RESTful API to auto
	// matically create a set or predefined rules, by default Sophos UTM will maintain those
	// sessions expecting additional API calls. You can close sessions if after creating your
	// rules you don’t anticipate subsequent API calls for the same process. At the last step,
	// you can set a header that will close the session.
	// X-Restd-Session: close
	//
	// Note – X-Restd-Session: close may cause a longer time for the next request.
	// Be sure to only send this command with the last request.
	//
	// Refer to Section 3.7 of Sophos documentation
	XRestdSession = "X-Restd-Session"

	Authoization = "Authorization"
)
