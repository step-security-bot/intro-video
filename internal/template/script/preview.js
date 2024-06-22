export default {
  setupIntroVideo: () => setupIntroVideo({ bubble, cta }),
  preload: (callback) => preload(videoConfig.url, callback),
  cleanUp: cleanUp,
}

