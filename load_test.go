// Copyright © 2020 Ettore Di Giacinto
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program; if not, see <http://www.gnu.org/licenses/>.

package extensions_test

import (
	"os"
	"path/filepath"

	. "github.com/mudler/cobra-extension"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Extension loader", func() {
	Context("a single extensions", func() {

		var extensions []ExtensionInterface
		BeforeEach(func() {
			extensions = Discover("project", "tests/ext-path")
		})

		It("Discovers it", func() {
			Expect(extensions).To(HaveLen(1))
		})

		It("Detect ShortName", func() {
			ext := extensions[0]
			Expect(ext.Short()).To(Equal("hello"), ext.String())
		})

		It("Detect correct Path", func() {
			ext := extensions[0]

			dir, err := os.Getwd()
			Expect(err).ToNot(HaveOccurred())

			Expect(ext.Path()).To(Equal(filepath.Join(dir, "tests", "ext-path", "project-hello")))
		})
	})
})
