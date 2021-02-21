package models

import "time"

type Context struct {
	Token                string `json:"token"`
	CurrentCustomerGroup struct {
		Name                                string      `json:"name"`
		DisplayGross                        bool        `json:"displayGross"`
		CustomFields                        interface{} `json:"customFields"`
		RegistrationActive                  bool        `json:"registrationActive"`
		RegistrationTitle                   interface{} `json:"registrationTitle"`
		RegistrationIntroduction            interface{} `json:"registrationIntroduction"`
		RegistrationOnlyCompanyRegistration interface{} `json:"registrationOnlyCompanyRegistration"`
		RegistrationSeoMetaDescription      interface{} `json:"registrationSeoMetaDescription"`
		UniqueIdentifier                    string      `json:"_uniqueIdentifier"`
		VersionID                           interface{} `json:"versionId"`
		Translated                          struct {
			Name                                string        `json:"name"`
			CustomFields                        []interface{} `json:"customFields"`
			RegistrationTitle                   interface{}   `json:"registrationTitle"`
			RegistrationIntroduction            interface{}   `json:"registrationIntroduction"`
			RegistrationOnlyCompanyRegistration interface{}   `json:"registrationOnlyCompanyRegistration"`
			RegistrationSeoMetaDescription      interface{}   `json:"registrationSeoMetaDescription"`
		} `json:"translated"`
		CreatedAt  time.Time   `json:"createdAt"`
		UpdatedAt  interface{} `json:"updatedAt"`
		Extensions struct {
			ForeignKeys struct {
				APIAlias string `json:"apiAlias"`
			} `json:"foreignKeys"`
		} `json:"extensions"`
		ID       string `json:"id"`
		APIAlias string `json:"apiAlias"`
	} `json:"currentCustomerGroup"`
	FallbackCustomerGroup struct {
		Name                                string      `json:"name"`
		DisplayGross                        bool        `json:"displayGross"`
		CustomFields                        interface{} `json:"customFields"`
		RegistrationActive                  bool        `json:"registrationActive"`
		RegistrationTitle                   interface{} `json:"registrationTitle"`
		RegistrationIntroduction            interface{} `json:"registrationIntroduction"`
		RegistrationOnlyCompanyRegistration interface{} `json:"registrationOnlyCompanyRegistration"`
		RegistrationSeoMetaDescription      interface{} `json:"registrationSeoMetaDescription"`
		UniqueIdentifier                    string      `json:"_uniqueIdentifier"`
		VersionID                           interface{} `json:"versionId"`
		Translated                          struct {
			Name                                string        `json:"name"`
			CustomFields                        []interface{} `json:"customFields"`
			RegistrationTitle                   interface{}   `json:"registrationTitle"`
			RegistrationIntroduction            interface{}   `json:"registrationIntroduction"`
			RegistrationOnlyCompanyRegistration interface{}   `json:"registrationOnlyCompanyRegistration"`
			RegistrationSeoMetaDescription      interface{}   `json:"registrationSeoMetaDescription"`
		} `json:"translated"`
		CreatedAt  time.Time   `json:"createdAt"`
		UpdatedAt  interface{} `json:"updatedAt"`
		Extensions struct {
			ForeignKeys struct {
				APIAlias string `json:"apiAlias"`
			} `json:"foreignKeys"`
		} `json:"extensions"`
		ID       string `json:"id"`
		APIAlias string `json:"apiAlias"`
	} `json:"fallbackCustomerGroup"`
	Currency struct {
		IsoCode              string      `json:"isoCode"`
		Factor               int         `json:"factor"`
		Symbol               string      `json:"symbol"`
		ShortName            string      `json:"shortName"`
		Name                 string      `json:"name"`
		Position             int         `json:"position"`
		DecimalPrecision     int         `json:"decimalPrecision"`
		CustomFields         interface{} `json:"customFields"`
		ShippingMethodPrices interface{} `json:"shippingMethodPrices"`
		IsSystemDefault      bool        `json:"isSystemDefault"`
		UniqueIdentifier     string      `json:"_uniqueIdentifier"`
		VersionID            interface{} `json:"versionId"`
		Translated           struct {
			ShortName    string        `json:"shortName"`
			Name         string        `json:"name"`
			CustomFields []interface{} `json:"customFields"`
		} `json:"translated"`
		CreatedAt  time.Time   `json:"createdAt"`
		UpdatedAt  interface{} `json:"updatedAt"`
		Extensions struct {
			InternalMappingStorage struct {
				APIAlias string `json:"apiAlias"`
			} `json:"internal_mapping_storage"`
			ForeignKeys struct {
				APIAlias string `json:"apiAlias"`
			} `json:"foreignKeys"`
		} `json:"extensions"`
		ID       string `json:"id"`
		APIAlias string `json:"apiAlias"`
	} `json:"currency"`
	SalesChannel struct {
		LanguageID              string      `json:"languageId"`
		CurrencyID              string      `json:"currencyId"`
		PaymentMethodID         string      `json:"paymentMethodId"`
		ShippingMethodID        string      `json:"shippingMethodId"`
		CountryID               string      `json:"countryId"`
		NavigationCategoryID    string      `json:"navigationCategoryId"`
		NavigationCategoryDepth int         `json:"navigationCategoryDepth"`
		FooterCategoryID        string      `json:"footerCategoryId"`
		ServiceCategoryID       interface{} `json:"serviceCategoryId"`
		Name                    string      `json:"name"`
		ShortName               interface{} `json:"shortName"`
		Configuration           interface{} `json:"configuration"`
		Active                  bool        `json:"active"`
		Maintenance             bool        `json:"maintenance"`
		TaxCalculationType      string      `json:"taxCalculationType"`
		Currency                struct {
			IsoCode              string      `json:"isoCode"`
			Factor               int         `json:"factor"`
			Symbol               string      `json:"symbol"`
			ShortName            string      `json:"shortName"`
			Name                 string      `json:"name"`
			Position             int         `json:"position"`
			DecimalPrecision     int         `json:"decimalPrecision"`
			CustomFields         interface{} `json:"customFields"`
			ShippingMethodPrices interface{} `json:"shippingMethodPrices"`
			IsSystemDefault      bool        `json:"isSystemDefault"`
			UniqueIdentifier     string      `json:"_uniqueIdentifier"`
			VersionID            interface{} `json:"versionId"`
			Translated           struct {
				ShortName    string        `json:"shortName"`
				Name         string        `json:"name"`
				CustomFields []interface{} `json:"customFields"`
			} `json:"translated"`
			CreatedAt  time.Time   `json:"createdAt"`
			UpdatedAt  interface{} `json:"updatedAt"`
			Extensions struct {
				InternalMappingStorage struct {
					APIAlias string `json:"apiAlias"`
				} `json:"internal_mapping_storage"`
				ForeignKeys struct {
					APIAlias string `json:"apiAlias"`
				} `json:"foreignKeys"`
			} `json:"extensions"`
			ID       string `json:"id"`
			APIAlias string `json:"apiAlias"`
		} `json:"currency"`
		Language       interface{} `json:"language"`
		PaymentMethod  interface{} `json:"paymentMethod"`
		ShippingMethod interface{} `json:"shippingMethod"`
		Country        interface{} `json:"country"`
		Domains        []struct {
			URL                         string        `json:"url"`
			CurrencyID                  string        `json:"currencyId"`
			Currency                    interface{}   `json:"currency"`
			SnippetSetID                string        `json:"snippetSetId"`
			SalesChannelID              string        `json:"salesChannelId"`
			LanguageID                  string        `json:"languageId"`
			Language                    interface{}   `json:"language"`
			CustomFields                interface{}   `json:"customFields"`
			SalesChannelDefaultHreflang interface{}   `json:"salesChannelDefaultHreflang"`
			HreflangUseOnlyLocale       bool          `json:"hreflangUseOnlyLocale"`
			UniqueIdentifier            string        `json:"_uniqueIdentifier"`
			VersionID                   interface{}   `json:"versionId"`
			Translated                  []interface{} `json:"translated"`
			CreatedAt                   time.Time     `json:"createdAt"`
			UpdatedAt                   interface{}   `json:"updatedAt"`
			Extensions                  struct {
				ForeignKeys struct {
					APIAlias string `json:"apiAlias"`
				} `json:"foreignKeys"`
			} `json:"extensions"`
			ID       string `json:"id"`
			APIAlias string `json:"apiAlias"`
		} `json:"domains"`
		CustomFields            interface{} `json:"customFields"`
		NavigationCategory      interface{} `json:"navigationCategory"`
		FooterCategory          interface{} `json:"footerCategory"`
		ServiceCategory         interface{} `json:"serviceCategory"`
		MailHeaderFooterID      interface{} `json:"mailHeaderFooterId"`
		CustomerGroupID         string      `json:"customerGroupId"`
		HreflangActive          bool        `json:"hreflangActive"`
		HreflangDefaultDomainID interface{} `json:"hreflangDefaultDomainId"`
		HreflangDefaultDomain   interface{} `json:"hreflangDefaultDomain"`
		LandingPages            interface{} `json:"landingPages"`
		UniqueIdentifier        string      `json:"_uniqueIdentifier"`
		VersionID               interface{} `json:"versionId"`
		Translated              struct {
			Name         string        `json:"name"`
			CustomFields []interface{} `json:"customFields"`
		} `json:"translated"`
		CreatedAt  time.Time `json:"createdAt"`
		UpdatedAt  time.Time `json:"updatedAt"`
		Extensions struct {
			ForeignKeys struct {
				APIAlias string `json:"apiAlias"`
			} `json:"foreignKeys"`
		} `json:"extensions"`
		ID                          string      `json:"id"`
		NavigationCategoryVersionID string      `json:"navigationCategoryVersionId"`
		FooterCategoryVersionID     interface{} `json:"footerCategoryVersionId"`
		ServiceCategoryVersionID    interface{} `json:"serviceCategoryVersionId"`
		APIAlias                    string      `json:"apiAlias"`
	} `json:"salesChannel"`
	TaxRules []struct {
		TaxRate          int           `json:"taxRate"`
		Name             string        `json:"name"`
		CustomFields     interface{}   `json:"customFields"`
		UniqueIdentifier string        `json:"_uniqueIdentifier"`
		VersionID        interface{}   `json:"versionId"`
		Translated       []interface{} `json:"translated"`
		CreatedAt        time.Time     `json:"createdAt"`
		UpdatedAt        interface{}   `json:"updatedAt"`
		Extensions       struct {
			ForeignKeys struct {
				APIAlias string `json:"apiAlias"`
			} `json:"foreignKeys"`
		} `json:"extensions"`
		ID       string `json:"id"`
		APIAlias string `json:"apiAlias"`
	} `json:"taxRules"`
	Customer      interface{} `json:"customer"`
	PaymentMethod struct {
		Name              string      `json:"name"`
		Description       string      `json:"description"`
		Position          int         `json:"position"`
		Active            bool        `json:"active"`
		AfterOrderEnabled bool        `json:"afterOrderEnabled"`
		Translations      interface{} `json:"translations"`
		MediaID           interface{} `json:"mediaId"`
		Media             interface{} `json:"media"`
		CustomFields      interface{} `json:"customFields"`
		ShortName         string      `json:"shortName"`
		UniqueIdentifier  string      `json:"_uniqueIdentifier"`
		VersionID         interface{} `json:"versionId"`
		Translated        struct {
			Name         string        `json:"name"`
			Description  string        `json:"description"`
			CustomFields []interface{} `json:"customFields"`
		} `json:"translated"`
		CreatedAt  time.Time   `json:"createdAt"`
		UpdatedAt  interface{} `json:"updatedAt"`
		Extensions struct {
			ForeignKeys struct {
				APIAlias string `json:"apiAlias"`
			} `json:"foreignKeys"`
		} `json:"extensions"`
		ID       string `json:"id"`
		APIAlias string `json:"apiAlias"`
	} `json:"paymentMethod"`
	ShippingMethod struct {
		Name           string      `json:"name"`
		Active         bool        `json:"active"`
		Description    interface{} `json:"description"`
		TrackingURL    interface{} `json:"trackingUrl"`
		DeliveryTimeID string      `json:"deliveryTimeId"`
		DeliveryTime   struct {
			Name             string      `json:"name"`
			Min              int         `json:"min"`
			Max              int         `json:"max"`
			Unit             string      `json:"unit"`
			CustomFields     interface{} `json:"customFields"`
			UniqueIdentifier string      `json:"_uniqueIdentifier"`
			VersionID        interface{} `json:"versionId"`
			Translated       struct {
				Name         string        `json:"name"`
				CustomFields []interface{} `json:"customFields"`
			} `json:"translated"`
			CreatedAt  time.Time   `json:"createdAt"`
			UpdatedAt  interface{} `json:"updatedAt"`
			Extensions struct {
				InternalMappingStorage struct {
					APIAlias string `json:"apiAlias"`
				} `json:"internal_mapping_storage"`
				ForeignKeys struct {
					APIAlias string `json:"apiAlias"`
				} `json:"foreignKeys"`
			} `json:"extensions"`
			ID       string `json:"id"`
			APIAlias string `json:"apiAlias"`
		} `json:"deliveryTime"`
		Translations interface{} `json:"translations"`
		CustomFields struct {
			SendcloudServicePointEnabled bool `json:"sendcloud_service_point_enabled"`
		} `json:"customFields"`
		AvailabilityRule interface{}   `json:"availabilityRule"`
		Prices           []interface{} `json:"prices"`
		MediaID          interface{}   `json:"mediaId"`
		Media            interface{}   `json:"media"`
		Tags             interface{}   `json:"tags"`
		TaxType          string        `json:"taxType"`
		Tax              interface{}   `json:"tax"`
		UniqueIdentifier string        `json:"_uniqueIdentifier"`
		VersionID        interface{}   `json:"versionId"`
		Translated       struct {
			Name         string `json:"name"`
			CustomFields struct {
				SendcloudServicePointEnabled bool `json:"sendcloud_service_point_enabled"`
			} `json:"customFields"`
			Description interface{} `json:"description"`
			TrackingURL interface{} `json:"trackingUrl"`
		} `json:"translated"`
		CreatedAt  time.Time `json:"createdAt"`
		UpdatedAt  time.Time `json:"updatedAt"`
		Extensions struct {
			ForeignKeys struct {
				APIAlias string `json:"apiAlias"`
			} `json:"foreignKeys"`
		} `json:"extensions"`
		ID       string `json:"id"`
		APIAlias string `json:"apiAlias"`
	} `json:"shippingMethod"`
	ShippingLocation struct {
		Country struct {
			Name                       string      `json:"name"`
			Iso                        string      `json:"iso"`
			Position                   int         `json:"position"`
			TaxFree                    bool        `json:"taxFree"`
			Active                     bool        `json:"active"`
			ShippingAvailable          bool        `json:"shippingAvailable"`
			Iso3                       string      `json:"iso3"`
			DisplayStateInRegistration bool        `json:"displayStateInRegistration"`
			ForceStateInRegistration   bool        `json:"forceStateInRegistration"`
			CompanyTaxFree             bool        `json:"companyTaxFree"`
			CheckVatIDPattern          bool        `json:"checkVatIdPattern"`
			VatIDPattern               string      `json:"vatIdPattern"`
			States                     interface{} `json:"states"`
			Translations               interface{} `json:"translations"`
			CustomFields               interface{} `json:"customFields"`
			UniqueIdentifier           string      `json:"_uniqueIdentifier"`
			VersionID                  interface{} `json:"versionId"`
			Translated                 struct {
				Name         string        `json:"name"`
				CustomFields []interface{} `json:"customFields"`
			} `json:"translated"`
			CreatedAt  time.Time   `json:"createdAt"`
			UpdatedAt  interface{} `json:"updatedAt"`
			Extensions struct {
				ForeignKeys struct {
					APIAlias string `json:"apiAlias"`
				} `json:"foreignKeys"`
			} `json:"extensions"`
			ID       string `json:"id"`
			APIAlias string `json:"apiAlias"`
		} `json:"country"`
		State    interface{} `json:"state"`
		Address  interface{} `json:"address"`
		APIAlias string      `json:"apiAlias"`
	} `json:"shippingLocation"`
	RulesIds         []string      `json:"rulesIds"`
	RulesLocked      bool          `json:"rulesLocked"`
	Permissions      []interface{} `json:"permissions"`
	PermisionsLocked bool          `json:"permisionsLocked"`
	APIAlias         string        `json:"apiAlias"`
}
