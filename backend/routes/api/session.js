const express = require('express')
const { setTokenCookie, restoreUser } = require('../../utils/auth');
const { User } = require('../../db/models');
const { check } = require('express-validator');
const { handleValidationErrors } = require('../../utils/validation');

const router = express.Router();

const validateLogin = [
    check('credential')
        .exists({ checkFalsy: true })
        .notEmpty()
        .withMessage('Email or username is required'),
    check('password')
        .exists({ checkFalsy: true })
        .withMessage('Password is required'),
    handleValidationErrors
];

const validateSignup = [
    check('email')
        .exists({ checkFalsy: true })
        .isEmail()
        .withMessage('Please provide a valid email.'),
    check('username')
        .exists({ checkFalsy: true })
        .isLength({ min: 4 })
        .withMessage('Username must be at least 4 characters long.'),
    check('username')
        .not()
        .isEmail()
        .withMessage('Username must not be an email.'),
    check('password')
        .exists({ checkFalsy: true })
        .isLength({ min: 6 })
        .withMessage('Password must be at least 6 characters long.'),
    handleValidationErrors
];

// Restore session
router.get(
    '/',
    restoreUser,
    (req, res) => {
        const { user } = req;
        if (user) {
            return res.json(user.toSafeObject());
        } else return res.json(null);
    }
);


// Log in
router.post(
    '/',
    validateLogin,
    async (req, res, next) => {
        const { credential, password } = req.body;
        const user = await User.login({ credential, password });
        if (!user) {
            const err = new Error('Invalid Credentials');
            err.status = 401;
            return next(err);
        }
        return res.json(user);
    }
);

router.post(
    '/new',
    validateSignup,
    async (req, res, next)=>{
        let user;
        try {
            User.signup(req.body).then(data => user = data)
        } catch {
            const err = new Error("Unable to complete signup!")
            err.status = 500;
            err.errors = {"uknown" : "Server failed to sign up new user"}
            return next(err)
        }
       res.json(user)
    }
)

// Log out
router.delete(
    '/',
    (_req, res) => {
        res.clearCookie('token');
        return res.json({ message: 'success' });
    }
);

module.exports = router;