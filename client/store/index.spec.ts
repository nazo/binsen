import assert from 'power-assert';
import { RootState, mutations } from './index';

it('test mutations setLang', () => {
  const state: RootState = new RootState();
  assert.equal(state.locale, 'en');
  mutations.setLang(state, 'ja');
  assert.equal(state.locale, 'ja');
});
