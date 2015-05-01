// Copyright 2011 <chaishushan@gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opencv

//#include "opencv.h"
//#cgo linux  pkg-config: opencv
//#cgo darwin pkg-config: opencv
//#cgo freebsd pkg-config: opencv
//#cgo windows LDFLAGS: -lopencv_core242.dll -lopencv_imgproc242.dll -lopencv_photo242.dll -lopencv_highgui242.dll -lstdc++
import "C"
import (
	//"errors"
	"unsafe"
)

func init() {
}

const (
	CV_BGR2GRAY  = C.CV_BGR2GRAY
	CV_BGR2BGRA  = C.CV_BGR2BGRA
	CV_RGBA2BGRA = C.CV_RGBA2BGRA
	CV_BGR2HSV   = C.CV_BGR2HSV

	CV_BLUR     = C.CV_BLUR
	CV_GAUSSIAN = C.CV_GAUSSIAN

	CV_8U  = C.CV_8U
	CV_8S  = C.CV_8S
	CV_16U = C.CV_16U
	CV_16S = C.CV_16S
	CV_32S = C.CV_32S
	CV_32F = C.CV_32F
	CV_64F = C.CV_64F

	CV_DXT_FORWARD = C.CV_DXT_FORWARD
	CV_DXT_INVERSE = C.CV_DXT_INVERSE
)

/* Smoothes array (removes noise) */
func Smooth(src, dst *IplImage, smoothtype,
	param1, param2 int, param3, param4 float64) {
	C.cvSmooth(unsafe.Pointer(src), unsafe.Pointer(dst), C.int(smoothtype),
		C.int(param1), C.int(param2), C.double(param3), C.double(param4),
	)
}

//CVAPI(void) cvSmooth( const CvArr* src, CvArr* dst,
//                      int smoothtype CV_DEFAULT(CV_GAUSSIAN),
//                      int param1 CV_DEFAULT(3),
//                      int param2 CV_DEFAULT(0),
//                      double param3 CV_DEFAULT(0),
//                      double param4 CV_DEFAULT(0));

/*
ConvertScale converts one image to another with optional linear transformation.
*/
func ConvertScale(src, dst *IplImage, scale, shift float64) {
	C.cvConvertScale(unsafe.Pointer(src), unsafe.Pointer(dst), C.double(scale), C.double(shift))
}

//CVAPI(void)  cvConvertScale( const CvArr* src,
//                             CvArr* dst,
//                             double scale CV_DEFAULT(1),
//                             double shift CV_DEFAULT(0) );

/* Converts input array pixels from one color space to another */
func CvtColor(src, dst *IplImage, code int) {
	C.cvCvtColor(unsafe.Pointer(src), unsafe.Pointer(dst), C.int(code))
}

//CVAPI(void)  cvCvtColor( const CvArr* src, CvArr* dst, int code );

/* Runs canny edge detector */
func Canny(image, edges *IplImage, threshold1, threshold2 float64, aperture_size int) {
	C.cvCanny(unsafe.Pointer(image), unsafe.Pointer(edges),
		C.double(threshold1), C.double(threshold2),
		C.int(aperture_size),
	)
}

//CVAPI(void)  cvCanny( const CvArr* image, CvArr* edges, double threshold1,
//                      double threshold2, int  aperture_size CV_DEFAULT(3) );

/* Calculates the first, second, third, or mixed image derivatives using an
* extended Sobel operator.  */
func Sobel(src, dst *IplImage, xorder, yorder, aperture_size int) {
	C.cvSobel(unsafe.Pointer(src), unsafe.Pointer(dst),
		C.int(xorder), C.int(yorder),
		C.int(aperture_size),
	)
}

// C: void cvSobel(const CvArr* src, CvArr* dst, int xorder, int yorder, int aperture_size=3 )

const (
	CV_INPAINT_NS    = C.CV_INPAINT_NS
	CV_INPAINT_TELEA = C.CV_INPAINT_TELEA
)

/* Inpaints the selected region in the image */
func Inpaint(src, inpaint_mask, dst *IplImage, inpaintRange float64, flags int) {
	C.cvInpaint(
		unsafe.Pointer(src),
		unsafe.Pointer(inpaint_mask),
		unsafe.Pointer(dst),
		C.double(inpaintRange),
		C.int(flags),
	)
}

//CVAPI(void) cvInpaint( const CvArr* src, const CvArr* inpaint_mask,
//                       CvArr* dst, double inpaintRange, int flags );

/* dst(idx) = lower <= src(idx) < upper */
func InRangeS(src *IplImage, lower Scalar, upper Scalar, dst *IplImage) {
	C.cvInRangeS(unsafe.Pointer(src), C.CvScalar(lower), C.CvScalar(upper), unsafe.Pointer(dst))
}

// CVAPI(void) cvInRangeS( const CvArr* src, CvScalar lower,
//                        CvScalar upper, CvArr* dst );

/* creates structuring element used for morphological operations */
func CreateStructuringElementEx(
	cols int, rows int,
	anchor_x int, anchor_y int,
	shape int, values []int) *IplConvKernel {
	return (*IplConvKernel)(C.cvCreateStructuringElementEx(
		C.int(cols), C.int(rows),
		C.int(anchor_x), C.int(anchor_y),
		C.int(shape), (*C.int)(unsafe.Pointer(&values))))
}

// CVAPI(IplConvKernel*)  cvCreateStructuringElementEx(
//             int cols, int  rows, int  anchor_x, int  anchor_y,
//             int shape, int* values CV_DEFAULT(NULL) );

/* releases structuring element */
func ReleaseStructuringElement(element **IplConvKernel) {
	C.cvReleaseStructuringElement((**C.IplConvKernel)(unsafe.Pointer(element)))
}

// CVAPI(void)  cvReleaseStructuringElement( IplConvKernel** element );

/* erodes input image (applies minimum filter) one or more times.
   If element pointer is NULL, 3x3 rectangular element is used */
func Erode(src, dst *IplImage, element *IplConvKernel, iterations int) {
	C.cvErode(unsafe.Pointer(src), unsafe.Pointer(dst), (*C.IplConvKernel)(unsafe.Pointer(element)), C.int(iterations))
}

// CVAPI(void)  cvErode( const CvArr* src, CvArr* dst,
//                       IplConvKernel* element CV_DEFAULT(NULL),
//                       int iterations CV_DEFAULT(1) );

/* dilates input image (applies maximum filter) one or more times.
   If element pointer is NULL, 3x3 rectangular element is used */
func Dilate(src, dst *IplImage, element *IplConvKernel, iterations int) {
	C.cvDilate(unsafe.Pointer(src), unsafe.Pointer(dst), (*C.IplConvKernel)(unsafe.Pointer(element)), C.int(iterations))
}

// CVAPI(void)  cvDilate( const CvArr* src, CvArr* dst,
//                        IplConvKernel* element CV_DEFAULT(NULL),
//                        int iterations CV_DEFAULT(1) );

/* dst(idx) = src1(idx) & src2(idx) */
func And(src1, src2, dst, mask *IplImage) {
	C.cvAnd(unsafe.Pointer(src1), unsafe.Pointer(src2), unsafe.Pointer(dst), unsafe.Pointer(mask))
}

// CVAPI(void) cvAnd( const CvArr* src1, const CvArr* src2,
//                   CvArr* dst, const CvArr* mask CV_DEFAULT(NULL));

/* Discrete Cosine Transform */
func DCT(src, dst *IplImage, flags int) {
	C.cvDCT(unsafe.Pointer(src), unsafe.Pointer(dst), C.int(flags))
}

// CVAPI(void)  cvDCT( const CvArr* src, CvArr* dst, int flags );

/* equalizes histogram of 8-bit single-channel image */
func EqualizeHist(src, dst *IplImage) {
	C.cvEqualizeHist(unsafe.Pointer(src), unsafe.Pointer(dst))
}

// CVAPI(void)  cvEqualizeHist( const CvArr* src, CvArr* dst );
/* Fills the connected component until the color difference gets large enough */
func FloodFill(image *IplImage, seed_point Point, new_val Scalar,
	low_diff Scalar, up_diff Scalar, comp *ConnectedComp, flags int, mask *IplImage) {
	C.cvFloodFill(unsafe.Pointer(image), C.cvPoint(C.int(seed_point.X), C.int(seed_point.Y)), C.CvScalar(new_val),
		C.CvScalar(low_diff), C.CvScalar(up_diff), (*C.CvConnectedComp)(unsafe.Pointer(comp)),
		C.int(flags), unsafe.Pointer(mask))
}

// CVAPI(void)  cvFloodFill( CvArr* image, CvPoint seed_point,
//                           CvScalar new_val, CvScalar lo_diff CV_DEFAULT(cvScalarAll(0)),
//                           CvScalar up_diff CV_DEFAULT(cvScalarAll(0)),
//                           CvConnectedComp* comp CV_DEFAULT(NULL),
//                           int flags CV_DEFAULT(4),
//                           CvArr* mask CV_DEFAULT(NULL));
