package customimagesearch

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// CustomInstanceClient is the the Bing Custom Image Search API lets you send an image search query to Bing and get
// back image search results customized to meet your custom search definition.
type CustomInstanceClient struct {
	BaseClient
}

// NewCustomInstanceClient creates an instance of the CustomInstanceClient client.
func NewCustomInstanceClient() CustomInstanceClient {
	return CustomInstanceClient{New()}
}

// ImageSearch sends the image search request.
// Parameters:
// customConfig - the identifier for the custom search configuration
// query - the user's search query term. The term cannot be empty. The term may contain [Bing Advanced
// Operators](http://msdn.microsoft.com/library/ff795620.aspx). For example, to limit images to a specific
// domain, use the [site:](http://msdn.microsoft.com/library/ff795613.aspx) operator. To help improve relevance
// of an insights query (see
// [insightsToken](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#insightstoken)),
// you should always include the user's query term. Use this parameter only with the Image Search API.Do not
// specify this parameter when calling the Trending Images API.
// acceptLanguage - a comma-delimited list of one or more languages to use for user interface strings. The list
// is in decreasing order of preference. For additional information, including expected format, see
// [RFC2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html). This header and the
// [setLang](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#setlang)
// query parameter are mutually exclusive; do not specify both. If you set this header, you must also specify
// the [cc](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#cc) query
// parameter. To determine the market to return results for, Bing uses the first supported language it finds
// from the list and combines it with the cc parameter value. If the list does not include a supported
// language, Bing finds the closest language and market that supports the request or it uses an aggregated or
// default market for the results. To determine the market that Bing used, see the BingAPIs-Market header. Use
// this header and the cc query parameter only if you specify multiple languages. Otherwise, use the
// [mkt](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#mkt) and
// [setLang](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#setlang)
// query parameters. A user interface string is a string that's used as a label in a user interface. There are
// few user interface strings in the JSON response objects. Any links to Bing.com properties in the response
// objects apply the specified language.
// userAgent - the user agent originating the request. Bing uses the user agent to provide mobile users with an
// optimized experience. Although optional, you are encouraged to always specify this header. The user-agent
// should be the same string that any commonly used browser sends. For information about user agents, see [RFC
// 2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html). The following are examples of user-agent
// strings. Windows Phone: Mozilla/5.0 (compatible; MSIE 10.0; Windows Phone 8.0; Trident/6.0; IEMobile/10.0;
// ARM; Touch; NOKIA; Lumia 822). Android: Mozilla / 5.0 (Linux; U; Android 2.3.5; en - us; SCH - I500 Build /
// GINGERBREAD) AppleWebKit / 533.1 (KHTML; like Gecko) Version / 4.0 Mobile Safari / 533.1. iPhone: Mozilla /
// 5.0 (iPhone; CPU iPhone OS 6_1 like Mac OS X) AppleWebKit / 536.26 (KHTML; like Gecko) Mobile / 10B142
// iPhone4; 1 BingWeb / 3.03.1428.20120423. PC: Mozilla / 5.0 (Windows NT 6.3; WOW64; Trident / 7.0; Touch;
// rv:11.0) like Gecko. iPad: Mozilla / 5.0 (iPad; CPU OS 7_0 like Mac OS X) AppleWebKit / 537.51.1 (KHTML,
// like Gecko) Version / 7.0 Mobile / 11A465 Safari / 9537.53
// clientID - bing uses this header to provide users with consistent behavior across Bing API calls. Bing often
// flights new features and improvements, and it uses the client ID as a key for assigning traffic on different
// flights. If you do not use the same client ID for a user across multiple requests, then Bing may assign the
// user to multiple conflicting flights. Being assigned to multiple conflicting flights can lead to an
// inconsistent user experience. For example, if the second request has a different flight assignment than the
// first, the experience may be unexpected. Also, Bing can use the client ID to tailor web results to that
// client ID’s search history, providing a richer experience for the user. Bing also uses this header to help
// improve result rankings by analyzing the activity generated by a client ID. The relevance improvements help
// with better quality of results delivered by Bing APIs and in turn enables higher click-through rates for the
// API consumer. IMPORTANT: Although optional, you should consider this header required. Persisting the client
// ID across multiple requests for the same end user and device combination enables 1) the API consumer to
// receive a consistent user experience, and 2) higher click-through rates via better quality of results from
// the Bing APIs. Each user that uses your application on the device must have a unique, Bing generated client
// ID. If you do not include this header in the request, Bing generates an ID and returns it in the
// X-MSEdge-ClientID response header. The only time that you should NOT include this header in a request is the
// first time the user uses your app on that device. Use the client ID for each Bing API request that your app
// makes for this user on the device. Persist the client ID. To persist the ID in a browser app, use a
// persistent HTTP cookie to ensure the ID is used across all sessions. Do not use a session cookie. For other
// apps such as mobile apps, use the device's persistent storage to persist the ID. The next time the user uses
// your app on that device, get the client ID that you persisted. Bing responses may or may not include this
// header. If the response includes this header, capture the client ID and use it for all subsequent Bing
// requests for the user on that device. If you include the X-MSEdge-ClientID, you must not include cookies in
// the request.
// clientIP - the IPv4 or IPv6 address of the client device. The IP address is used to discover the user's
// location. Bing uses the location information to determine safe search behavior. Although optional, you are
// encouraged to always specify this header and the X-Search-Location header. Do not obfuscate the address (for
// example, by changing the last octet to 0). Obfuscating the address results in the location not being
// anywhere near the device's actual location, which may result in Bing serving erroneous results.
// location - a semicolon-delimited list of key/value pairs that describe the client's geographical location.
// Bing uses the location information to determine safe search behavior and to return relevant local content.
// Specify the key/value pair as <key>:<value>. The following are the keys that you use to specify the user's
// location. lat (required): The latitude of the client's location, in degrees. The latitude must be greater
// than or equal to -90.0 and less than or equal to +90.0. Negative values indicate southern latitudes and
// positive values indicate northern latitudes. long (required): The longitude of the client's location, in
// degrees. The longitude must be greater than or equal to -180.0 and less than or equal to +180.0. Negative
// values indicate western longitudes and positive values indicate eastern longitudes. re (required): The
// radius, in meters, which specifies the horizontal accuracy of the coordinates. Pass the value returned by
// the device's location service. Typical values might be 22m for GPS/Wi-Fi, 380m for cell tower triangulation,
// and 18,000m for reverse IP lookup. ts (optional): The UTC UNIX timestamp of when the client was at the
// location. (The UNIX timestamp is the number of seconds since January 1, 1970.) head (optional): The client's
// relative heading or direction of travel. Specify the direction of travel as degrees from 0 through 360,
// counting clockwise relative to true north. Specify this key only if the sp key is nonzero. sp (optional):
// The horizontal velocity (speed), in meters per second, that the client device is traveling. alt (optional):
// The altitude of the client device, in meters. are (optional): The radius, in meters, that specifies the
// vertical accuracy of the coordinates. Specify this key only if you specify the alt key. Although many of the
// keys are optional, the more information that you provide, the more accurate the location results are.
// Although optional, you are encouraged to always specify the user's geographical location. Providing the
// location is especially important if the client's IP address does not accurately reflect the user's physical
// location (for example, if the client uses VPN). For optimal results, you should include this header and the
// X-MSEdge-ClientIP header, but at a minimum, you should include this header.
// aspect - filter images by the following aspect ratios. All: Do not filter by aspect.Specifying this value is
// the same as not specifying the aspect parameter. Square: Return images with standard aspect ratio. Wide:
// Return images with wide screen aspect ratio. Tall: Return images with tall aspect ratio.
// colorParameter - filter images by the following color options. ColorOnly: Return color images. Monochrome:
// Return black and white images. Return images with one of the following dominant colors: Black, Blue, Brown,
// Gray, Green, Orange, Pink, Purple, Red, Teal, White, Yellow
// countryCode - a 2-character country code of the country where the results come from. For a list of possible
// values, see [Market
// Codes](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#market-codes).
// If you set this parameter, you must also specify the
// [Accept-Language](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#acceptlanguage)
// header. Bing uses the first supported language it finds from the languages list, and combine that language
// with the country code that you specify to determine the market to return results for. If the languages list
// does not include a supported language, Bing finds the closest language and market that supports the request,
// or it may use an aggregated or default market for the results instead of a specified one. You should use
// this query parameter and the Accept-Language query parameter only if you specify multiple languages;
// otherwise, you should use the mkt and setLang query parameters. This parameter and the
// [mkt](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#mkt) query
// parameter are mutually exclusive—do not specify both.
// count - the number of images to return in the response. The actual number delivered may be less than
// requested. The default is 35. The maximum value is 150. You use this parameter along with the offset
// parameter to page results.For example, if your user interface displays 20 images per page, set count to 20
// and offset to 0 to get the first page of results.For each subsequent page, increment offset by 20 (for
// example, 0, 20, 40). Use this parameter only with the Image Search API.Do not specify this parameter when
// calling the Insights, Trending Images, or Web Search APIs.
// freshness - filter images by the following discovery options. Day: Return images discovered by Bing within
// the last 24 hours. Week: Return images discovered by Bing within the last 7 days. Month: Return images
// discovered by Bing within the last 30 days.
// height - filter images that have the specified height, in pixels. You may use this filter with the size
// filter to return small images that have a height of 150 pixels.
// ID - an ID that uniquely identifies an image. Use this parameter to ensure that the specified image is the
// first image in the list of images that Bing returns. The
// [Image](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#image)
// object's imageId field contains the ID that you set this parameter to.
// imageContent - filter images by the following content types. Face: Return images that show only a person's
// face. Portrait: Return images that show only a person's head and shoulders.
// imageType - filter images by the following image types. AnimatedGif: Return only animated GIFs. Clipart:
// Return only clip art images. Line: Return only line drawings. Photo: Return only photographs(excluding line
// drawings, animated Gifs, and clip art). Shopping: Return only images that contain items where Bing knows of
// a merchant that is selling the items. This option is valid in the en - US market only.Transparent: Return
// only images with a transparent background.
// license - filter images by the following license types. All: Do not filter by license type.Specifying this
// value is the same as not specifying the license parameter. Any: Return images that are under any license
// type. The response doesn't include images that do not specify a license or the license is unknown. Public:
// Return images where the creator has waived their exclusive rights, to the fullest extent allowed by law.
// Share: Return images that may be shared with others. Changing or editing the image might not be allowed.
// Also, modifying, sharing, and using the image for commercial purposes might not be allowed. Typically, this
// option returns the most images. ShareCommercially: Return images that may be shared with others for personal
// or commercial purposes. Changing or editing the image might not be allowed. Modify: Return images that may
// be modified, shared, and used. Changing or editing the image might not be allowed. Modifying, sharing, and
// using the image for commercial purposes might not be allowed. ModifyCommercially: Return images that may be
// modified, shared, and used for personal or commercial purposes. Typically, this option returns the fewest
// images. For more information about these license types, see [Filter Images By License
// Type](http://go.microsoft.com/fwlink/?LinkId=309768).
// market - the market where the results come from. Typically, mkt is the country where the user is making the
// request from. However, it could be a different country if the user is not located in a country where Bing
// delivers results. The market must be in the form <language code>-<country code>. For example, en-US. The
// string is case insensitive. For a list of possible market values, see [Market
// Codes](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#market-codes).
// NOTE: If known, you are encouraged to always specify the market. Specifying the market helps Bing route the
// request and return an appropriate and optimal response. If you specify a market that is not listed in
// [Market
// Codes](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#market-codes),
// Bing uses a best fit market code based on an internal mapping that is subject to change. This parameter and
// the [cc](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#cc) query
// parameter are mutually exclusive—do not specify both.
// maxFileSize - filter images that are less than or equal to the specified file size. The maximum file size
// that you may specify is 520,192 bytes. If you specify a larger value, the API uses 520,192. It is possible
// that the response may include images that are slightly larger than the specified maximum. You may specify
// this filter and minFileSize to filter images within a range of file sizes.
// maxHeight - filter images that have a height that is less than or equal to the specified height. Specify the
// height in pixels. You may specify this filter and minHeight to filter images within a range of heights. This
// filter and the height filter are mutually exclusive.
// maxWidth - filter images that have a width that is less than or equal to the specified width. Specify the
// width in pixels. You may specify this filter and maxWidth to filter images within a range of widths. This
// filter and the width filter are mutually exclusive.
// minFileSize - filter images that are greater than or equal to the specified file size. The maximum file size
// that you may specify is 520,192 bytes. If you specify a larger value, the API uses 520,192. It is possible
// that the response may include images that are slightly smaller than the specified minimum. You may specify
// this filter and maxFileSize to filter images within a range of file sizes.
// minHeight - filter images that have a height that is greater than or equal to the specified height. Specify
// the height in pixels. You may specify this filter and maxHeight to filter images within a range of heights.
// This filter and the height filter are mutually exclusive.
// minWidth - filter images that have a width that is greater than or equal to the specified width. Specify the
// width in pixels. You may specify this filter and maxWidth to filter images within a range of widths. This
// filter and the width filter are mutually exclusive.
// offset - the zero-based offset that indicates the number of images to skip before returning images. The
// default is 0. The offset should be less than
// ([totalEstimatedMatches](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#totalestimatedmatches)
// - count). Use this parameter along with the count parameter to page results. For example, if your user
// interface displays 20 images per page, set count to 20 and offset to 0 to get the first page of results. For
// each subsequent page, increment offset by 20 (for example, 0, 20, 40). It is possible for multiple pages to
// include some overlap in results. To prevent duplicates, see
// [nextOffset](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#nextoffset).
// Use this parameter only with the Image API. Do not specify this parameter when calling the Trending Images
// API or the Web Search API.
// safeSearch - filter images for adult content. The following are the possible filter values. Off: May return
// images with adult content. If the request is through the Image Search API, the response includes thumbnail
// images that are clear (non-fuzzy). However, if the request is through the Web Search API, the response
// includes thumbnail images that are pixelated (fuzzy). Moderate: If the request is through the Image Search
// API, the response doesn't include images with adult content. If the request is through the Web Search API,
// the response may include images with adult content (the thumbnail images are pixelated (fuzzy)). Strict: Do
// not return images with adult content. The default is Moderate. If the request comes from a market that
// Bing's adult policy requires that safeSearch is set to Strict, Bing ignores the safeSearch value and uses
// Strict. If you use the site: query operator, there is the chance that the response may contain adult content
// regardless of what the safeSearch query parameter is set to. Use site: only if you are aware of the content
// on the site and your scenario supports the possibility of adult content.
// size - filter images by the following sizes. All: Do not filter by size. Specifying this value is the same
// as not specifying the size parameter. Small: Return images that are less than 200x200 pixels. Medium: Return
// images that are greater than or equal to 200x200 pixels but less than 500x500 pixels. Large: Return images
// that are 500x500 pixels or larger. Wallpaper: Return wallpaper images. You may use this parameter along with
// the height or width parameters. For example, you may use height and size to request small images that are
// 150 pixels tall.
// setLang - the language to use for user interface strings. Specify the language using the ISO 639-1 2-letter
// language code. For example, the language code for English is EN. The default is EN (English). Although
// optional, you should always specify the language. Typically, you set setLang to the same language specified
// by mkt unless the user wants the user interface strings displayed in a different language. This parameter
// and the
// [Accept-Language](https://docs.microsoft.com/en-us/rest/api/cognitiveservices/bing-images-api-v7-reference#acceptlanguage)
// header are mutually exclusive; do not specify both. A user interface string is a string that's used as a
// label in a user interface. There are few user interface strings in the JSON response objects. Also, any
// links to Bing.com properties in the response objects apply the specified language.
// width - filter images that have the specified width, in pixels. You may use this filter with the size filter
// to return small images that have a width of 150 pixels.
func (client CustomInstanceClient) ImageSearch(ctx context.Context, customConfig string, query string, acceptLanguage string, userAgent string, clientID string, clientIP string, location string, aspect ImageAspect, colorParameter ImageColor, countryCode string, count *int32, freshness Freshness, height *int32, ID string, imageContent ImageContent, imageType ImageType, license ImageLicense, market string, maxFileSize *int64, maxHeight *int64, maxWidth *int64, minFileSize *int64, minHeight *int64, minWidth *int64, offset *int64, safeSearch SafeSearch, size ImageSize, setLang string, width *int32) (result Images, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomInstanceClient.ImageSearch")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ImageSearchPreparer(ctx, customConfig, query, acceptLanguage, userAgent, clientID, clientIP, location, aspect, colorParameter, countryCode, count, freshness, height, ID, imageContent, imageType, license, market, maxFileSize, maxHeight, maxWidth, minFileSize, minHeight, minWidth, offset, safeSearch, size, setLang, width)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customimagesearch.CustomInstanceClient", "ImageSearch", nil, "Failure preparing request")
		return
	}

	resp, err := client.ImageSearchSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "customimagesearch.CustomInstanceClient", "ImageSearch", resp, "Failure sending request")
		return
	}

	result, err = client.ImageSearchResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "customimagesearch.CustomInstanceClient", "ImageSearch", resp, "Failure responding to request")
	}

	return
}

// ImageSearchPreparer prepares the ImageSearch request.
func (client CustomInstanceClient) ImageSearchPreparer(ctx context.Context, customConfig string, query string, acceptLanguage string, userAgent string, clientID string, clientIP string, location string, aspect ImageAspect, colorParameter ImageColor, countryCode string, count *int32, freshness Freshness, height *int32, ID string, imageContent ImageContent, imageType ImageType, license ImageLicense, market string, maxFileSize *int64, maxHeight *int64, maxWidth *int64, minFileSize *int64, minHeight *int64, minWidth *int64, offset *int64, safeSearch SafeSearch, size ImageSize, setLang string, width *int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"Endpoint": client.Endpoint,
	}

	queryParameters := map[string]interface{}{
		"customConfig": autorest.Encode("query", customConfig),
		"q":            autorest.Encode("query", query),
	}
	if len(string(aspect)) > 0 {
		queryParameters["aspect"] = autorest.Encode("query", aspect)
	}
	if len(string(colorParameter)) > 0 {
		queryParameters["color"] = autorest.Encode("query", colorParameter)
	}
	if len(countryCode) > 0 {
		queryParameters["cc"] = autorest.Encode("query", countryCode)
	}
	if count != nil {
		queryParameters["count"] = autorest.Encode("query", *count)
	}
	if len(string(freshness)) > 0 {
		queryParameters["freshness"] = autorest.Encode("query", freshness)
	}
	if height != nil {
		queryParameters["height"] = autorest.Encode("query", *height)
	}
	if len(ID) > 0 {
		queryParameters["id"] = autorest.Encode("query", ID)
	}
	if len(string(imageContent)) > 0 {
		queryParameters["imageContent"] = autorest.Encode("query", imageContent)
	}
	if len(string(imageType)) > 0 {
		queryParameters["imageType"] = autorest.Encode("query", imageType)
	}
	if len(string(license)) > 0 {
		queryParameters["license"] = autorest.Encode("query", license)
	}
	if len(market) > 0 {
		queryParameters["mkt"] = autorest.Encode("query", market)
	}
	if maxFileSize != nil {
		queryParameters["maxFileSize"] = autorest.Encode("query", *maxFileSize)
	}
	if maxHeight != nil {
		queryParameters["maxHeight"] = autorest.Encode("query", *maxHeight)
	}
	if maxWidth != nil {
		queryParameters["maxWidth"] = autorest.Encode("query", *maxWidth)
	}
	if minFileSize != nil {
		queryParameters["minFileSize"] = autorest.Encode("query", *minFileSize)
	}
	if minHeight != nil {
		queryParameters["minHeight"] = autorest.Encode("query", *minHeight)
	}
	if minWidth != nil {
		queryParameters["minWidth"] = autorest.Encode("query", *minWidth)
	}
	if offset != nil {
		queryParameters["offset"] = autorest.Encode("query", *offset)
	}
	if len(string(safeSearch)) > 0 {
		queryParameters["safeSearch"] = autorest.Encode("query", safeSearch)
	}
	if len(string(size)) > 0 {
		queryParameters["size"] = autorest.Encode("query", size)
	}
	if len(setLang) > 0 {
		queryParameters["setLang"] = autorest.Encode("query", setLang)
	}
	if width != nil {
		queryParameters["width"] = autorest.Encode("query", *width)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{Endpoint}/bingcustomsearch/v7.0", urlParameters),
		autorest.WithPath("/images/search"),
		autorest.WithQueryParameters(queryParameters),
		autorest.WithHeader("X-BingApis-SDK", "true"))
	if len(acceptLanguage) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("Accept-Language", autorest.String(acceptLanguage)))
	}
	if len(userAgent) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("User-Agent", autorest.String(userAgent)))
	}
	if len(clientID) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-MSEdge-ClientID", autorest.String(clientID)))
	}
	if len(clientIP) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-MSEdge-ClientIP", autorest.String(clientIP)))
	}
	if len(location) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("X-Search-Location", autorest.String(location)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ImageSearchSender sends the ImageSearch request. The method will close the
// http.Response Body if it receives an error.
func (client CustomInstanceClient) ImageSearchSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ImageSearchResponder handles the response to the ImageSearch request. The method always
// closes the http.Response Body.
func (client CustomInstanceClient) ImageSearchResponder(resp *http.Response) (result Images, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
