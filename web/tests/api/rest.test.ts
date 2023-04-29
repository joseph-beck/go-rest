import { get, post } from '../../src/api/rest';
import { expect } from 'chai';
import 'mocha';

describe('Options tests', () => {
    it('checking default options', async () => {
        const data = await get();
        expect(data).to.be.not.null;
    });
});

describe('Options tests', () => {
    it('checking default options', async () => {
        const preSend = await get();
        expect(preSend).to.be.not.null;

        post('testdata', 'this is a piece of test data from ts');
        
        const postSend = await get();
        expect(postSend).to.be.not.null;

        let preSendLength: number = preSend!.length
        let postSendLength: number = postSend!.length

        expect(postSendLength).greaterThan(preSendLength);
    });
});