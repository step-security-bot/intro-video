/**
* @returns {void}
*/
function run() {
  const videoConfig = getVideoConfig();

  if (!videoConfig.url) {
    console.error('No video URL provided');
    return;
  }

  preload(videoConfig, () => {
    const bubble = getBubble();
    const cta = getCTA();
    setupIntroVideo({ videoConfig, bubble, cta });
  });
}

return {
  run,
  cleanup,
}

