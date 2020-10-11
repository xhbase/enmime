package enmime

import (
	"errors"
	"net/textproto"
	"time"

	"github.com/jhillyerd/enmime/internal/stringutil"
)

// Build performs some basic validations, then constructs a tree of Part structs from the configured
// MailBuilder.  It will set the Date header to now if it was not explicitly set.
func (p MailBuilder) BuildDraft() (*Part, error) {
	if p.err != nil {
		return nil, p.err
	}
	// Validations
	if p.from.Address == "" {
		return nil, errors.New("from not set")
	}
	if p.subject == "" {
		return nil, errors.New("subject not set")
	}
	if len(p.to)+len(p.cc)+len(p.bcc) == 0 {
		return nil, errors.New("no recipients (to, cc, bcc) set")
	}
	// Fully loaded structure; the presence of text, html, inlines, and attachments will determine
	// how much is necessary:
	//
	//  multipart/mixed
	//  |- multipart/related
	//  |  |- multipart/alternative
	//  |  |  |- text/plain
	//  |  |  `- text/html
	//  |  `- inlines..
	//  `- attachments..
	//
	// We build this tree starting at the leaves, re-rooting as needed.
	var root, part *Part
	if p.text != nil || p.html == nil {
		root = NewPart(ctTextPlain)
		root.Content = p.text
		root.Charset = utf8
	}
	if p.html != nil {
		part = NewPart(ctTextHTML)
		part.Content = p.html
		part.Charset = utf8
		if root == nil {
			root = part
		} else {
			root.NextSibling = part
		}
	}
	if p.text != nil && p.html != nil {
		// Wrap Text & HTML bodies
		part = root
		root = NewPart(ctMultipartAltern)
		root.AddChild(part)
	}
	if len(p.inlines) > 0 {
		part = root
		root = NewPart(ctMultipartRelated)
		root.AddChild(part)
		for _, ip := range p.inlines {
			// Copy inline Part to isolate mutations
			part = &Part{}
			*part = *ip
			part.Header = make(textproto.MIMEHeader)
			root.AddChild(part)
		}
	}
	if len(p.attachments) > 0 {
		part = root
		root = NewPart(ctMultipartMixed)
		root.AddChild(part)
		for _, ap := range p.attachments {
			// Copy attachment Part to isolate mutations
			part = &Part{}
			*part = *ap
			part.Header = make(textproto.MIMEHeader)
			root.AddChild(part)
		}
	}
	// Headers
	h := root.Header
	h.Set(hnMIMEVersion, "1.0")
	h.Set("From", p.from.String())
	h.Set("Subject", p.subject)
	if len(p.to) > 0 {
		h.Set("To", stringutil.JoinAddress(p.to))
	}
	if len(p.cc) > 0 {
		h.Set("Cc", stringutil.JoinAddress(p.cc))
	}
	if p.replyTo.Address != "" {
		h.Set("Reply-To", p.replyTo.String())
	}
	date := p.date
	if date.IsZero() {
		date = time.Now()
	}
	h.Set("Date", date.Format(time.RFC1123Z))
	for k, v := range p.header {
		for _, s := range v {
			h.Add(k, s)
		}
	}
	return root, nil
}
