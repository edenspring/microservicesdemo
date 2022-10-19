const router = require('express').Router();
const sessionRouter = require('./session')

router.use('session', sessionRouter)

router.post('/test', function (req, res) {
    res.json({ requestBody: req.body });
});

module.exports = router;