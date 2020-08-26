package plist

// AppTransportSecurity allows you to specify a description of changes made to
// the default security for HTTP connections.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/NSAppTransportSecurity for more information.
type AppTransportSecurity struct {
	allowArbitraryLoads                bool
	allowArbitraryLoadsMedia           bool
	allowArbitraryLoadsWebContent      bool
	allowArbitraryLoadsLocalNetworking bool

	exceptionDomains map[string]*ATSExceptionDomain
}

// Apply will apply AppTransportSecurity against the specified PropertyList.
func (s *AppTransportSecurity) Apply(p *PropertyList) {
	p.Set(keyNSAppTransportSecurity, s.build())
}

// AllowArbitraryLoads specifies a boolean value indicating whether App
// Transport Security restrictions are disabled for all network connections.
// In iOS 10 and later and macOS 10.12 and later, the value of the
// NSAllowsArbitraryLoads key is ignored—and the default value of NO used
// instead—if any of the following keys are present in your app’s Information Property List file:
//  NSAllowsArbitraryLoadsForMedia
//  NSAllowsArbitraryLoadsInWebContent
//  NSAllowsLocalNetworking
// See https://developer.apple.com/documentation/bundleresources/information_property_list/nsapptransportsecurity/nsallowsarbitraryloads for more information.
func (s *AppTransportSecurity) AllowArbitraryLoads(v bool) {
	s.allowArbitraryLoads = v
}

// AllowArbitraryLoadForMedia specifies a boolean value indicating whether all App
// Transport Security restrictions are disabled for requests made using the AV
// Foundation framework.
// https://developer.apple.com/documentation/bundleresources/information_property_list/nsapptransportsecurity/nsallowsarbitraryloadsformedia for more information.
func (s *AppTransportSecurity) AllowArbitraryLoadForMedia(v bool) {
	s.allowArbitraryLoadsMedia = v
}

// AllowArbitraryLoadForWebContent specifies a boolean value indicating whether
// all App Transport Security restrictions are disabled for requests made from
// web views.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/nsapptransportsecurity/NSAllowsArbitraryLoadsInWebContent for more information.
func (s *AppTransportSecurity) AllowArbitraryLoadForWebContent(v bool) {
	s.allowArbitraryLoadsWebContent = v
}

// AllowArbitraryLoadForLocalNetworking specifies a boolean value indicating
// whether to allow loading of local resources.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/nsapptransportsecurity/NSAllowsLocalNetworking for more information.
func (s *AppTransportSecurity) AllowArbitraryLoadForLocalNetworking(v bool) {
	s.allowArbitraryLoadsLocalNetworking = v
}

// ExceptionDomain allows you tp specify a custom configuration for  App
// Transport Security named domains.
// See https://developer.apple.com/documentation/bundleresources/information_property_list/nsapptransportsecurity/nsexceptiondomains for more information.
func (s *AppTransportSecurity) ExceptionDomain(domain string, f func(d *ATSExceptionDomain)) {
	if s.exceptionDomains == nil {
		s.exceptionDomains = map[string]*ATSExceptionDomain{}
	}

	d := &ATSExceptionDomain{}
	d.init()

	s.exceptionDomains[domain] = d
	f(s.exceptionDomains[domain])
}

func (s *AppTransportSecurity) build() map[string]interface{} {
	data := map[string]interface{}{}

	data[atsNSAllowsArbitraryLoads] = s.allowArbitraryLoads

	if s.allowArbitraryLoadsMedia {
		data[atsNSAllowsArbitraryLoadsForMedia] = s.allowArbitraryLoadsMedia
	}

	if s.allowArbitraryLoadsWebContent {
		data[atsNSAllowsArbitraryLoadsInWebContent] = s.allowArbitraryLoadsWebContent
	}

	if s.allowArbitraryLoadsLocalNetworking {
		data[atsNSAllowsLocalNetworking] = s.allowArbitraryLoadsLocalNetworking
	}

	if len(s.exceptionDomains) > 0 {
		exceptions := map[string]interface{}{}
		for domain, exception := range s.exceptionDomains {
			exceptions[domain] = exception.build()
		}

		data[atsNSExceptionDomains] = exceptions
	}

	return data
}

func _canary() {
	ats := &AppTransportSecurity{}
	ats.AllowArbitraryLoads(true)
	ats.AllowArbitraryLoadForMedia(true)
	ats.AllowArbitraryLoadForWebContent(true)
	ats.AllowArbitraryLoadForLocalNetworking(true)
	ats.ExceptionDomain("google.com", func(d *ATSExceptionDomain) {
		d.IncludesSubdomains(true)
		d.AllowsInsecureHTTPLoads(true)
		d.MinimumTLSVersion("1.2")
		d.RequiresForwardSecrecy(true)
		d.RequiresCertificateTransparency(true)
	})
}
