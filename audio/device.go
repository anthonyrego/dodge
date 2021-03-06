package audio

/*
#ifdef _GOSMF_OSX_
  #include <CoreFoundation/CoreFoundation.h>
#endif

#include <OpenAL/al.h>
#include <OpenAL/alc.h>

ALCcontext *context;
ALCdevice *device;

void openAudioDevice() {
	const ALCchar *default_device;

	default_device = alcGetString(NULL, ALC_DEFAULT_DEVICE_SPECIFIER);
	if ((device = alcOpenDevice(default_device)) == NULL)
	{
		fprintf(stderr, "failed to open sound device\n");
		return;
	}
	context = alcCreateContext(device, NULL);

	alcMakeContextCurrent(context);

	alcProcessContext(context);

	if (alcGetError(device) != ALC_NO_ERROR) {
		fprintf(stderr, "Failed to open audio device\n");
	}
}

void closeAudioDevice()
{
	alcMakeContextCurrent(NULL);
	alcDestroyContext(context);
	alcCloseDevice(device);
}
*/
import "C"

func initDevice() {
	C.openAudioDevice()
}

func destroyDevice() {
	C.closeAudioDevice()
}
