package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/streams/vocab"
)

func main() {

	// Create a Note object
	var note vocab.ActivityStreamsNote = streams.NewActivityStreamsNote()

	// Create an `id` property and set it on the Note
	id, _ := url.Parse("https://example.com/some/path/to/this/note")
	var idProperty vocab.JSONLDIdProperty = streams.NewJSONLDIdProperty()
	idProperty.Set(id)

	// Set the `id` property on our Note.
	note.SetJSONLDId(idProperty)

	// Let's try to add content to our note. First, let's get the property.
	contentProperty := note.GetActivityStreamsContent()

	// WARNING: Missing properties are `nil`!
	if contentProperty == nil {
		// Create a new property and set it on the note.
		contentProperty = streams.NewActivityStreamsContentProperty()
		// Treat properties as pointers, not values. Setting a
		// property is not a value-copy so if we modify
		// the property later, any modification will be
		// reflected in the note.
		note.SetActivityStreamsContent(contentProperty)
	}

	// Now we are guaranteed a non-`nil` property: let's add content!
	contentProperty.AppendXMLSchemaString("jorts")

	// The "published" property is functional: It can only have at most one value.
	published := streams.NewActivityStreamsPublishedProperty()
	// We can set a time...
	published.Set(time.Now())
	// ...or, in this very unusual practice, set it as an IRI
	// iri, _ := url.Parse("https://go-fed.org/some/path")
	// published.SetIRI(iri)

	if published.IsIRI() {
		fmt.Println(published.GetIRI())
	} else if published.IsXMLSchemaDateTime() {
		fmt.Println(published.Get())
	}

	// The "object" property is non-functional: It can have many values.
	object := streams.NewActivityStreamsObjectProperty()
	// We can append...
	object.AppendActivityStreamsNote(note)
	// ...or prepend...
	object.PrependActivityStreamsArticle(streams.NewActivityStreamsArticle())
	// ...and IRIs too
	iri, _ := url.Parse("https://go-fed.org/foo")
	object.AppendIRI(iri)

	// for iter := object.Begin(); iter != object.End(); iter = iter.Next() {
	// 	fmt.Println(iter.GetIRI())
	// 	if iter.IsActivityStreamsNote() {
	// 		note := iter.GetActivityStreamsNote()
	// 		fmt.Println(note.GetJSONLDId())
	// 	} else if iter.IsActivityStreamsArticle() {
	// 		article := iter.GetActivityStreamsArticle()
	// 		fmt.Println(article.GetJSONLDId())
	// 	} else if iter.IsIRI() {
	// 		iri := iter.GetIRI()
	// 		fmt.Println(iri.String())
	// 	}
	// }

	// Deserialize a JSON payload
	jsonstr := `{
		"@context": [
		  "https://www.w3.org/ns/activitystreams",
		  "https://w3id.org/security/v1",
		  {
			"Curve25519Key": "toot:Curve25519Key",
			"Device": "toot:Device",
			"Ed25519Key": "toot:Ed25519Key",
			"Ed25519Signature": "toot:Ed25519Signature",
			"EncryptedMessage": "toot:EncryptedMessage",
			"IdentityProof": "toot:IdentityProof",
			"PropertyValue": "schema:PropertyValue",
			"alsoKnownAs": {
			  "@id": "as:alsoKnownAs",
			  "@type": "@id"
			},
			"cipherText": "toot:cipherText",
			"claim": {
			  "@id": "toot:claim",
			  "@type": "@id"
			},
			"deviceId": "toot:deviceId",
			"devices": {
			  "@id": "toot:devices",
			  "@type": "@id"
			},
			"discoverable": "toot:discoverable",
			"featured": {
			  "@id": "toot:featured",
			  "@type": "@id"
			},
			"featuredTags": {
			  "@id": "toot:featuredTags",
			  "@type": "@id"
			},
			"fingerprintKey": {
			  "@id": "toot:fingerprintKey",
			  "@type": "@id"
			},
			"focalPoint": {
			  "@container": "@list",
			  "@id": "toot:focalPoint"
			},
			"identityKey": {
			  "@id": "toot:identityKey",
			  "@type": "@id"
			},
			"manuallyApprovesFollowers": "as:manuallyApprovesFollowers",
			"messageFranking": "toot:messageFranking",
			"messageType": "toot:messageType",
			"movedTo": {
			  "@id": "as:movedTo",
			  "@type": "@id"
			},
			"publicKeyBase64": "toot:publicKeyBase64",
			"schema": "http://schema.org#",
			"suspended": "toot:suspended",
			"toot": "http://joinmastodon.org/ns#",
			"value": "schema:value"
		  }
		],
		"alsoKnownAs": [
		  "https://tooting.ai/users/Gargron"
		],
		"attachment": [
		  {
			"name": "Patreon",
			"type": "PropertyValue",
			"value": "<a href=\"https://www.patreon.com/mastodon\" rel=\"me nofollow noopener noreferrer\" target=\"_blank\"><span class=\"invisible\">https://www.</span><span class=\"\">patreon.com/mastodon</span><span class=\"invisible\"></span></a>"
		  },
		  {
			"name": "Homepage",
			"type": "PropertyValue",
			"value": "<a href=\"https://zeonfederated.com\" rel=\"me nofollow noopener noreferrer\" target=\"_blank\"><span class=\"invisible\">https://</span><span class=\"\">zeonfederated.com</span><span class=\"invisible\"></span></a>"
		  },
		  {
			"name": "gargron",
			"signatureAlgorithm": "keybase",
			"signatureValue": "5cfc20c7018f2beefb42a68836da59a792e55daa4d118498c9b1898de7e845690f",
			"type": "IdentityProof"
		  }
		],
		"devices": "https://mastodon.social/users/Gargron/collections/devices",
		"discoverable": true,
		"endpoints": {
		  "sharedInbox": "https://mastodon.social/inbox"
		},
		"featured": "https://mastodon.social/users/Gargron/collections/featured",
		"featuredTags": "https://mastodon.social/users/Gargron/collections/tags",
		"followers": "https://mastodon.social/users/Gargron/followers",
		"following": "https://mastodon.social/users/Gargron/following",
		"icon": {
		  "mediaType": "image/jpeg",
		  "type": "Image",
		  "url": "https://files.mastodon.social/accounts/avatars/000/000/001/original/d96d39a0abb45b92.jpg"
		},
		"id": "https://mastodon.social/users/Gargron",
		"image": {
		  "mediaType": "image/png",
		  "type": "Image",
		  "url": "https://files.mastodon.social/accounts/headers/000/000/001/original/c91b871f294ea63e.png"
		},
		"inbox": "https://mastodon.social/users/Gargron/inbox",
		"manuallyApprovesFollowers": false,
		"name": "Eugen",
		"outbox": "https://mastodon.social/users/Gargron/outbox",
		"preferredUsername": "Gargron",
		"publicKey": {
		  "id": "https://mastodon.social/users/Gargron#main-key",
		  "owner": "https://mastodon.social/users/Gargron",
		  "publicKeyPem": "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvXc4vkECU2/CeuSo1wtn\nFoim94Ne1jBMYxTZ9wm2YTdJq1oiZKif06I2fOqDzY/4q/S9uccrE9Bkajv1dnkO\nVm31QjWlhVpSKynVxEWjVBO5Ienue8gND0xvHIuXf87o61poqjEoepvsQFElA5ym\novljWGSA/jpj7ozygUZhCXtaS2W5AD5tnBQUpcO0lhItYPYTjnmzcc4y2NbJV8hz\n2s2G8qKv8fyimE23gY1XrPJg+cRF+g4PqFXujjlJ7MihD9oqtLGxbu7o1cifTn3x\nBfIdPythWu5b4cujNsB3m3awJjVmx+MHQ9SugkSIYXV0Ina77cTNS0M2PYiH1PFR\nTwIDAQAB\n-----ENDPUBLIC KEY-----\n"
		},
		"published": "2016-03-16T00:00:00Z",
		"summary": "<p>Developer of Mastodon and administrator of mastodon.social. I post service announcements, development updates, and personal stuff.</p>",
		"tag": [],
		"type": "Person",
		"url": "https://mastodon.social/@Gargron"
	  }`
	// jsonstr := `{
	// 	"@context": "https://www.w3.org/ns/activitystreams",
	// 	"id":       "https://go-fed.org/foo",
	// 	"name":     "Foo Bar",
	// 	"inbox":    "https://go-fed.org/foo/inbox",
	// 	"outbox":   "https://go-fed.org/foo/outbox",
	// 	"type":     "Person",
	// 	"url":      "https://go-fed.org/foo"
	// }`

	var m map[string]interface{}
	_ = json.Unmarshal([]byte(jsonstr), &m)

	// Next, we prepare a streams.JSONResolver, providing one or more callbacks.
	var person vocab.ActivityStreamsPerson
	// var ib vocab.ActivityStreamsObject
	resolver, _ := streams.NewJSONResolver(
		func(c context.Context, p vocab.ActivityStreamsPerson) error {
			// Store the person in the enclosing scope, for later.
			person = p
			return nil
		},
		// func(c context.Context, i vocab.ActivityStreamsObject) error {
		// 	// Store the person in the enclosing scope, for later.
		// 	ib = i
		// 	return nil
		// },
		func(c context.Context, note vocab.ActivityStreamsNote) error {
			// We can treat the type differently.
			fmt.Println(note)
			return nil
		},
	)
	// It will call back a function we provide if it is of a matching type,
	// or returns streams.ErrNoCallbackMatch when we didn't give it a matcher for
	// the type, or streams.ErrUnhandledType if it is a type unknown to Go-Fed.
	ctx := context.Background()
	_ = resolver.Resolve(ctx, m)

	// Serialize to a JSON payload
	var jsonmap map[string]interface{}
	jsonmap, _ = streams.Serialize(person) // WARNING: Do not call the Serialize() method on person
	b, _ := json.Marshal(jsonmap)
	fmt.Println(string(b))

	outbox := person.GetActivityStreamsOutbox()
	fmt.Println(outbox.GetIRI())
	if outbox.IsActivityStreamsOrderedCollection() {
		fmt.Println("IsActivityStreamsOrderedCollection")
	} else if outbox.IsActivityStreamsOrderedCollectionPage() {
		fmt.Println("IsActivityStreamsOrderedCollectionPage")
	} else if outbox.IsIRI() {
		fmt.Println("IsIRI")
	} else {
		fmt.Println("IsNeither")
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", outbox.GetIRI().String(), nil)
	fmt.Println("GETting " + outbox.GetIRI().String())
	req.Header.Add("Accept", "application/activity+json")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	// page := oc.GetActivityStreamsCurrent().GetActivityStreamsOrderedCollectionPage()
	// orderedItems := page.GetActivityStreamsOrderedItems()
	// for iter := orderedItems.Begin(); iter != orderedItems.End(); iter = iter.Next() {
	// 	fmt.Println(iter.GetIRI())
	// }
}
