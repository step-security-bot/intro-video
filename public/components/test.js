import { LitElement, html, css, unsafeCSS} from 'https://cdn.jsdelivr.net/gh/lit/dist@3/core/lit-core.min.js';

import globalStyles from '/style.css' with { type: 'css' };

class Test extends LitElement {
  static styles = css`
    ${unsafeCSS([...globalStyles.rules].map(rule => rule.cssText).join(''))}
    :host {
      display: block
    }
  `;

  static properties = {
    counter: { type: Number }
  };

  constructor() {
    super();
    this.counter = 0;
  }

  render() {
    return html`
      <div class="p-4 bg-gray-200">
        <div>Hello from Lit!</div>
        <div>Counter: ${this.counter}</div>
        <button @click="${this.increment}">Increment</button>
      </div>
    `;
  }

  increment() {
    this.counter++;
  }
}

customElements.define('test-abc', Test);

