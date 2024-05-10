/**
* @param {string} videoUrl
* @returns {void}
*/
export default function run(videoUrl) {
  preload(videoUrl);

  const initialScrollPosition = window.scrollY;

  function handleScroll() {
    let scrollPosition = window.scrollY;
    if (Math.abs(initialScrollPosition - scrollPosition) > 100) {
      setupIntroVideo();
      window.removeEventListener('scroll', handleScroll);
    }
  }

  window.addEventListener('scroll', handleScroll);

}

/**
* @param {string} videoUrl
* @returns {void}
*/
function preload(videoUrl) {
  const container = document.querySelector('#intro-video');

  const video = document.createElement('video');

  video.id = 'intro-video-player';
  video.classList.add('iv-player');

  video.muted = true;
  video.loop = true;
  video.draggable = false;
  video.src = videoUrl;

  container.appendChild(video);
}

function createBubble() {
  const bubble = document.createElement('div');
  bubble.id = 'bubble-text';
  bubble.classList.add('iv-bubble');
  bubble.textContent = bubbleConfig?.textContent ?? 'Hello!';
  return bubble;
}

function createCta() {
  const cta = document.createElement('button');
  cta.classList.add('iv-cta-button');
  cta.textContent = ctaConfig?.textContent ?? 'Message Me';
  return cta;
}

function setupIntroVideo() {
  const container = document.querySelector('#intro-video');

  const card = document.createElement('div');
  card.classList.add('iv-card');

  let bubble = null;
  if (configBubble && configBubble.enabled) {
    bubble = createBubble();
  }

  let cta = null;
  if (configCta && configCta.enabled) {
    cta = createCta();
  }

  const videoWrapper = document.createElement('div');
  videoWrapper.classList.add('iv-player-wrapper');

  /**
  * @type {HTMLVideoElement}
  */
  const video = document.querySelector('#intro-video-player');
  video.style.display = 'block';

  const progressBar = document.createElement('progress');
  progressBar.id = 'intro-video-progressbar';
  progressBar.classList.add('iv-progressbar');
  progressBar.value = 0;
  progressBar.max = 100;

  const button = document.createElement('button');
  button.classList.add('iv-close-button');
  button.innerHTML = '&times;';

  video.addEventListener('timeupdate', function() {
    const percentage = (video.currentTime / video.duration) * 100;
    progressBar.value = percentage;
  });

  button.onclick = () => {
    card.style.opacity = 0;
    setTimeout(() => {
      card.remove();
    }, 500);
  }


  videoWrapper.onclick = () => {
    card.classList.add('large');
    video.muted = false;
    videoWrapper.appendChild(cta);
    if (bubble) {
      bubble.remove();
    }
  }

  videoWrapper.appendChild(video);
  videoWrapper.appendChild(progressBar);
  card.appendChild(videoWrapper);
  card.appendChild(button);
  card.appendChild(bubble);
  container.appendChild(card);
  video.play();
}
