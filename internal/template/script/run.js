/**
* @returns {void}
*/
function run() {
  const videoConfig = getVideoConfig();

  if (!videoConfig.url) {
    console.error('No video URL provided');
    return;
  }

  function handleScroll() {
    let scrollPosition = window.scrollY;
    if (Math.abs(initialScrollPosition - scrollPosition) > 100 && videoConfig.small.height > 0) {
      const bubble = getBubble();
      const cta = getCTA();
      setupIntroVideo({ videoConfig, bubble, cta });
      window.removeEventListener('scroll', handleScroll);
    }
  }

  const initialScrollPosition = window.scrollY;
  preload(videoConfig, () => {
    window.addEventListener('scroll', handleScroll);
  });
}

run();

