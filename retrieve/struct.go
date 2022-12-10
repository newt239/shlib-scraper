package retrieve

import (
	"encoding/xml"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Book struct {
	Name         string                     `dynamodbav:"name" json:"name"`
	Bid          int                        `dynamodbav:"bid" json:"bid"`
	Title        string                     `dynamodbav:"title" json:"title"`
	Author       string                     `dynamodbav:"author" json:"author"`
	Publisher    string                     `dynamodbav:"publisher" json:"publisher"`
	Pubdate      string                     `dynamodbav:"pubdate" json:"pubdate"`
	Lanove       bool                       `dynamodbav:"lanove" json:"lanove"`
	Tameshiyomi  bool                       `dynamodbav:"tameshiyomi" json:"tameshiyomi"`
	Isbn         string                     `dynamodbav:"isbn" json:"isbn"`
	Status       string                     `dynamodbav:"status" json:"status"`
	LocalCount   int                        `dynamodbav:"localCount" json:"localCount"`
	ReserveCount int                        `dynamodbav:"reserveCount" json:"reserveCount"`
	Source       string                     `dynamodbav:"source" json:"source"`
	CreatedAt    dynamodbattribute.UnixTime `dynamodbav:"created_at" json:"created_at"`
}

type BookWord struct {
	Name      string                     `dynamodbav:"name" json:"name"`
	Bid       int                        `dynamodbav:"bid" json:"bid"`
	Text      []string                   `dynamodbav:"text" json:"text"`
	CreatedAt dynamodbattribute.UnixTime `dynamodbav:"created_at" json:"created_at"`
}

type BookInfo struct {
	Name        string `dynamodbav:"name" json:"name"`
	Title       string `dynamodbav:"title" json:"title"`
	Author      string `dynamodbav:"author" json:"author"`
	Publisher   string `dynamodbav:"publisher" json:"publisher"`
	Pubdate     string `dynamodbav:"pubdate" json:"pubdate"`
	Lanove      bool   `dynamodbav:"lanove" json:"lanove"`
	Tameshiyomi bool   `dynamodbav:"tameshiyomi" json:"tameshiyomi"`
	Isbn        string `dynamodbav:"isbn" json:"isbn"`
	Source      string `dynamodbav:"source" json:"source"`
}

type OpenBDResponse []struct {
	Onix struct {
		RecordReference   string `json:"RecordReference"`
		NotificationType  string `json:"NotificationType"`
		ProductIdentifier struct {
			ProductIDType string `json:"ProductIDType"`
			IDValue       string `json:"IDValue"`
		} `json:"ProductIdentifier"`
		DescriptiveDetail struct {
			ProductComposition string `json:"ProductComposition"`
			ProductForm        string `json:"ProductForm"`
			ProductFormDetail  string `json:"ProductFormDetail"`
			TitleDetail        struct {
				TitleType    string `json:"TitleType"`
				TitleElement struct {
					TitleElementLevel string `json:"TitleElementLevel"`
					TitleText         struct {
						Collationkey string `json:"collationkey"`
						Content      string `json:"content"`
					} `json:"TitleText"`
					Subtitle struct {
						Collationkey string `json:"collationkey"`
						Content      string `json:"content"`
					} `json:"Subtitle"`
				} `json:"TitleElement"`
			} `json:"TitleDetail"`
			Contributor []struct {
				SequenceNumber  string   `json:"SequenceNumber"`
				ContributorRole []string `json:"ContributorRole"`
				PersonName      struct {
					Collationkey string `json:"collationkey"`
					Content      string `json:"content"`
				} `json:"PersonName"`
				BiographicalNote string `json:"BiographicalNote,omitempty"`
			} `json:"Contributor"`
			Language []struct {
				LanguageRole string `json:"LanguageRole"`
				LanguageCode string `json:"LanguageCode"`
				CountryCode  string `json:"CountryCode"`
			} `json:"Language"`
			Extent []struct {
				ExtentType  string `json:"ExtentType"`
				ExtentValue string `json:"ExtentValue"`
				ExtentUnit  string `json:"ExtentUnit"`
			} `json:"Extent"`
			Subject []struct {
				MainSubject             string `json:"MainSubject,omitempty"`
				SubjectSchemeIdentifier string `json:"SubjectSchemeIdentifier"`
				SubjectCode             string `json:"SubjectCode"`
			} `json:"Subject"`
		} `json:"DescriptiveDetail"`
		CollateralDetail struct {
			TextContent        []CollateralDetailTextContent
			SupportingResource []struct {
				ResourceContentType string `json:"ResourceContentType"`
				ContentAudience     string `json:"ContentAudience"`
				ResourceMode        string `json:"ResourceMode"`
				ResourceVersion     []struct {
					ResourceForm           string `json:"ResourceForm"`
					ResourceVersionFeature []struct {
						ResourceVersionFeatureType string `json:"ResourceVersionFeatureType"`
						FeatureValue               string `json:"FeatureValue"`
					} `json:"ResourceVersionFeature"`
					ResourceLink string `json:"ResourceLink"`
				} `json:"ResourceVersion"`
			} `json:"SupportingResource"`
		} `json:"CollateralDetail"`
		PublishingDetail struct {
			Imprint struct {
				ImprintIdentifier []struct {
					ImprintIDType string `json:"ImprintIDType"`
					IDValue       string `json:"IDValue"`
				} `json:"ImprintIdentifier"`
				ImprintName string `json:"ImprintName"`
			} `json:"Imprint"`
			PublishingDate []struct {
				PublishingDateRole string `json:"PublishingDateRole"`
				Date               string `json:"Date"`
			} `json:"PublishingDate"`
		} `json:"PublishingDetail"`
		ProductSupply struct {
			MarketPublishingDetail struct {
				MarketPublishingStatus     string `json:"MarketPublishingStatus"`
				MarketPublishingStatusNote string `json:"MarketPublishingStatusNote"`
			} `json:"MarketPublishingDetail"`
			SupplyDetail struct {
				ProductAvailability string `json:"ProductAvailability"`
				Price               []struct {
					PriceType    string `json:"PriceType"`
					PriceAmount  string `json:"PriceAmount"`
					CurrencyCode string `json:"CurrencyCode"`
				} `json:"Price"`
			} `json:"SupplyDetail"`
		} `json:"ProductSupply"`
	} `json:"onix"`
	Hanmoto struct {
		Toji             string `json:"toji"`
		Zaiko            int    `json:"zaiko"`
		Maegakinado      string `json:"maegakinado"`
		Kaisetsu105W     string `json:"kaisetsu105w"`
		Tsuiki           string `json:"tsuiki"`
		Genrecodetrc     int    `json:"genrecodetrc"`
		Ndccode          string `json:"ndccode"`
		Kankoukeitai     string `json:"kankoukeitai"`
		Sonotatokkijikou string `json:"sonotatokkijikou"`
		Jushoujouhou     string `json:"jushoujouhou"`
		Furokusonota     string `json:"furokusonota"`
		Dokushakakikomi  string `json:"dokushakakikomi"`
		Zasshicode       string `json:"zasshicode"`
		Jyuhan           []struct {
			Date    string `json:"date"`
			Ctime   string `json:"ctime"`
			Suri    int    `json:"suri"`
			Comment string `json:"comment,omitempty"`
		} `json:"jyuhan"`
		Hatsubai       string `json:"hatsubai"`
		Hatsubaiyomi   string `json:"hatsubaiyomi"`
		Lanove         bool   `json:"lanove"`
		Hastameshiyomi bool   `json:"hastameshiyomi"`
		Author         []struct {
			Listseq     int    `json:"listseq"`
			Dokujikubun string `json:"dokujikubun"`
		} `json:"author"`
		Datemodified  string `json:"datemodified"`
		Datecreated   string `json:"datecreated"`
		Kanrenshoisbn string `json:"kanrenshoisbn"`
		Reviews       []struct {
			PostUser   string `json:"post_user"`
			Reviewer   string `json:"reviewer"`
			SourceID   int    `json:"source_id"`
			KubunID    int    `json:"kubun_id"`
			Source     string `json:"source"`
			Choyukan   string `json:"choyukan"`
			Han        string `json:"han"`
			Link       string `json:"link"`
			Appearance string `json:"appearance"`
			Gou        string `json:"gou"`
		} `json:"reviews"`
		Hanmotoinfo struct {
			Name     string `json:"name"`
			Yomi     string `json:"yomi"`
			URL      string `json:"url"`
			Twitter  string `json:"twitter"`
			Facebook string `json:"facebook"`
		} `json:"hanmotoinfo"`
		Dateshuppan string `json:"dateshuppan"`
	} `json:"hanmoto"`
	Summary struct {
		Isbn      string `json:"isbn"`
		Title     string `json:"title"`
		Volume    string `json:"volume"`
		Series    string `json:"series"`
		Publisher string `json:"publisher"`
		Pubdate   string `json:"pubdate"`
		Cover     string `json:"cover"`
		Author    string `json:"author"`
	} `json:"summary"`
}
type CollateralDetailTextContent struct {
	TextType        string `json:"TextType"`
	ContentAudience string `json:"ContentAudience"`
	Text            string `json:"Text"`
}
type NDLApiResponse struct {
	XMLName            xml.Name `xml:"searchRetrieveResponse"`
	Text               string   `xml:",chardata"`
	Xmlns              string   `xml:"xmlns,attr"`
	Version            string   `xml:"version"`
	NumberOfRecords    string   `xml:"numberOfRecords"`
	NextRecordPosition string   `xml:"nextRecordPosition"`
	ExtraResponseData  struct {
		Text   string `xml:",chardata"`
		Facets struct {
			Text string `xml:",chardata"`
			Lst  []struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
				Int  []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"int"`
			} `xml:"lst"`
		} `xml:"facets"`
	} `xml:"extraResponseData"`
	Records struct {
		Text   string `xml:",chardata"`
		Record []struct {
			Text          string `xml:",chardata"`
			RecordSchema  string `xml:"recordSchema"`
			RecordPacking string `xml:"recordPacking"`
			RecordData    struct {
				Text string `xml:",chardata"`
				Dc   struct {
					Text           string   `xml:",chardata"`
					Dc             string   `xml:"dc,attr"`
					SrwDc          string   `xml:"srw_dc,attr"`
					Xsi            string   `xml:"xsi,attr"`
					SchemaLocation string   `xml:"schemaLocation,attr"`
					Title          string   `xml:"title"`
					Creator        string   `xml:"creator"`
					Description    []string `xml:"description"`
					Publisher      string   `xml:"publisher"`
					Language       string   `xml:"language"`
				} `xml:"dc"`
			} `xml:"recordData"`
			RecordPosition string `xml:"recordPosition"`
		} `xml:"record"`
	} `xml:"records"`
}

type YahooMAApiResponse struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		Tokens [][]string `json:"tokens"`
	} `json:"result"`
}
