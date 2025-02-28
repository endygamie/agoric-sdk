import { E } from '@endo/eventual-send';
import { Far } from '@endo/marshal';

console.log(`=> loading bootstrap.js`);

export function buildRootObject(_vatPowers) {
  return Far('root', {
    bootstrap(vats) {
      const pa = E(vats.bob).genPromise1();
      const pb = E(vats.bob).genPromise2();
      E(vats.bob).usePromises([pa], [pb]);
    },
  });
}
