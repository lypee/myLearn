package pdf

import (
	"bytes"
	 	"html/template"
	"strings"

	logger "git.100tal.com/wangxiao_go_lib/xesLogger"
		)

//pdf requestpdf struct
type RequestPdf struct {
	body string
}

//new request to pdf function
func NewRequestPdf(body string) *RequestPdf {
	return &RequestPdf{
		body: body,
	}
}

//parsing template function
func (r *RequestPdf) ParseTemplate(templateFileName string, data interface{}) error {

	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}

//generate pdf function
func (r *RequestPdf) GeneratePDF(htmlBody string) (*wkhtmltopdf.PDFGenerator, bool, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		logger.E("NewPDFGenerator", "err:[%+v]", err)
		return pdfg, false, err
	}

	reader := strings.NewReader(htmlBody)
	pdfg.AddPage(wkhtmltopdf.NewPageReader(reader))

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		logger.E("GeneratePDF", "err:[%+v]", err)
		return pdfg, false, err
	}
	return pdfg, true, nil
}
