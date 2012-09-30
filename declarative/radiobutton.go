// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package declarative

import (
	"github.com/lxn/walk"
)

type RadioButton struct {
	AssignTo      **walk.RadioButton
	Name          string
	MinSize       Size
	MaxSize       Size
	StretchFactor int
	Row           int
	RowSpan       int
	Column        int
	ColumnSpan    int
	ContextMenu   Menu
	Font          Font
	Text          string
	OnClicked     walk.EventHandler
}

func (rb RadioButton) Create(parent walk.Container) error {
	w, err := walk.NewRadioButton(parent)
	if err != nil {
		return err
	}

	return InitWidget(rb, w, func() error {
		if err := w.SetText(rb.Text); err != nil {
			return err
		}

		if rb.OnClicked != nil {
			w.Clicked().Attach(rb.OnClicked)
		}

		if rb.AssignTo != nil {
			*rb.AssignTo = w
		}

		return nil
	})
}

func (rb RadioButton) CommonInfo() (name string, minSize, maxSize Size, stretchFactor, row, rowSpan, column, columnSpan int, contextMenu *Menu) {
	return rb.Name, rb.MinSize, rb.MaxSize, rb.StretchFactor, rb.Row, rb.RowSpan, rb.Column, rb.ColumnSpan, &rb.ContextMenu
}

func (rb RadioButton) Font_() *Font {
	return &rb.Font
}
