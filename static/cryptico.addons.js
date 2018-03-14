// global deserializer method
fromJSON = function(key) {
	let json = JSON.parse(key);
	if (json.type !== 'RSAKey') return null;
	let rsa = new RSAKey();
	rsa.setPrivateEx(json.n, json.e, json.d, json.p, json.q, json.dmp1, json.dmq1, json.coeff);
	return rsa;
}
// instance serializer, JSON.stringify(object_with_RSAKey) will serialize as expected.
RSAKey.prototype.toJSON = function() {
	return JSON.stringify({
		type: 'RSAKey',
		coeff: this.coeff.toString(16),
		d: this.d.toString(16),
		dmp1: this.dmp1.toString(16),
		dmq1: this.dmq1.toString(16),
		e: this.e.toString(16),
		n: this.n.toString(16),
		p: this.p.toString(16),
		q: this.q.toString(16)
	})
}
