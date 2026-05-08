import {showMessage} from "../dialog/message";
import {getCloudURL} from "../config/util/about";
import {isDisabledFeature} from "../protyle/util/compatibility";

const isSyncProviderMembershipCheckBypassed = () => isDisabledFeature("sync-provider-membership-check-bypass");

export const needSubscribe = (tip = window.siyuan.languages._kernel[29]) => {
    if (isSyncProviderMembershipCheckBypassed()) {
        return false;
    }

    if (window.siyuan.user && (window.siyuan.user.userSiYuanProExpireTime === -1 || window.siyuan.user.userSiYuanProExpireTime > 0)) {
        return false;
    }
    if (tip) {
        if (tip === window.siyuan.languages._kernel[29] && window.siyuan.config.system.container === "ios") {
            showMessage(window.siyuan.languages._kernel[122]);
        } else {
            if (tip === window.siyuan.languages._kernel[29]) {
                tip = window.siyuan.languages._kernel[29].replaceAll("${accountServer}", getCloudURL(""));
            }
            showMessage(tip);
        }
    }
    return true;
};

export const isPaidUser = () => {
    if (isSyncProviderMembershipCheckBypassed()) {
        return true;
    }
    return window.siyuan.user && (0 === window.siyuan.user.userSiYuanSubscriptionStatus || 1 === window.siyuan.user.userSiYuanOneTimePayStatus);
};
