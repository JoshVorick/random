#include "opencv2/opencv.hpp"
#include "opencv2/nonfree/features2d.hpp"

#include <vector>

using namespace cv;
using namespace std;

int main(int argc, const char *argv[]) {
	CascadeClassifier haar_cascade;
	haar_cascade.load("haarcascade_frontalface_alt.xml");
	VideoCapture capCam(0);
	Mat frame;

	capCam >> frame;
	while(frame.data) {

		Mat gray;
		cvtColor(frame, gray, CV_BGR2GRAY);
		vector< Rect_<int> > faces;
		haar_cascade.detectMultiScale(gray, faces);

		for (int i=0; i<faces.size(); i++) {
			rectangle(frame, faces[i], CV_RGB(255,100,0), 1);
		}

		imshow("OpenGL", frame);
		char key = (char) waitKey(20);
		if (key == 27 || key == 'q')
			break;
		capCam >> frame;
	}
	return 0;
}
