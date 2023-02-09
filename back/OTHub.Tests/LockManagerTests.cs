using System;
using System.Threading.Tasks;
using OTHub.BackendSync;
using Xunit;

namespace OTHub.Tests
{
    public class LockManagerTests
    {

        [Fact]
        public async Task LockManager_ConfirmLockingIsSingleEntry()
        {
            Assert.False(LockManager.IsInUse(LockType.OfferCreated));

            using (await LockManager.Lock(LockType.OfferCreated))
            {
                Assert.True(LockManager.IsInUse(LockType.OfferCreated));
            }

            Assert.False(LockManager.IsInUse(LockType.OfferCreated));

            var type = typeof(Exception);
            using (await LockManager.Lock(LockType.OfferCreated))
            {
                await Assert.ThrowsAsync(type, async () => await LockManager.Lock(LockType.OfferCreated, 0));
            }
        }

    }
}