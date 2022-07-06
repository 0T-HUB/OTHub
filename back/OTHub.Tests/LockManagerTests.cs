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
            LockRelease lck = LockManager.GetLock(LockType.OfferCreated);

            Assert.Equal(1, lck.LockCurrentCount);

            await lck.Lock();

            Assert.Equal(0, lck.LockCurrentCount);

            var type = typeof(Exception);
            await Assert.ThrowsAsync(type, () => lck.Lock(0));
        } 

    }
}